import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom' // Importa esto
import './index.css'
import CryptoApp from './App/CryptoApp'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        {/* Si el usuario entra a la raíz "/", mándalo a BTC por defecto */}
        <Route path="/" element={<Navigate to="/BTCUSDT" />} />
        
        {/* Esta ruta captura lo que pongas después del "/" y lo mete en 'symbol' */}
        <Route path="/:symbol" element={<CryptoApp />} />
      </Routes>
    </BrowserRouter>
  </StrictMode>,
)