import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Activity, ArrowUpRight, Globe } from "lucide-react"; // Iconos pro
import { Badge } from "@/components/ui/badge";
import type { CryptoData } from "@/interfaces/interfaceCrypto";

export default function App() {
  const { symbol } = useParams<{ symbol: string }>();
  const [data, setData] = useState<CryptoData | null>(null);
  const [prevPrice, setPrevPrice] = useState<number>(0);

  useEffect(() => {
    // 2. Cada vez que el 'symbol' en la URL cambie, este efecto se ejecuta
    const fetchData = async () => {
      if (!symbol) return;

      // Forzamos mayúsculas para Binance
      const cleanSymbol = symbol.toUpperCase();

      try {
        const response = await fetch(
          `http://localhost:8081/binance/tracking/${cleanSymbol}`,
        );

        const result = await response.json();
        setData(result.data || result);
        // GUARDAR EL PRECIO ANTERIOR ANTES DE ACTUALIZAR EL NUEVO
        setData((prevData) => {
          if (prevData) setPrevPrice(parseFloat(prevData.price));
          return result || result;
        });
      } catch (error) {
        console.error("Error:", error);
      }
    };

    fetchData();

    // 3. Opcional: Intervalo que respeta el símbolo actual de la URL
    const interval = setInterval(fetchData, 500);
    return () => clearInterval(interval);
  }, [symbol]); // <--- IMPORTANTE: 'symbol' debe estar aquí para que el efecto se reinicie al cambiar la URL
  const priceColor = !data
    ? "text-white"
    : parseFloat(data.price) >= prevPrice
      ? "text-green-400"
      : "text-red-400";

  return (
    <div className="min-h-screen bg-[#09090b] text-zinc-100 selection:bg-blue-500/30">
      {/* Background Decorativo: Un resplandor radial muy tenue */}
      <div className="fixed inset-0 bg-[radial-gradient(circle_farthest-side_at_50%_-20%,rgba(59,130,246,0.1),transparent)] pointer-events-none" />

      <main className="relative z-10 max-w-5xl mx-auto pt-20 px-6">
        <div className="flex flex-col md:flex-row gap-6 items-start">
          {/* Card Principal de Precio */}
          <Card className="flex-1 bg-zinc-950/50 border-zinc-800 backdrop-blur-xl shadow-2xl">
            <CardHeader className="space-y-1">
              <div className="flex items-center justify-between">
                <Badge
                  variant="outline"
                  className="bg-blue-500/10 text-blue-400 border-blue-500/20 gap-1 uppercase tracking-widest text-[10px]"
                >
                  <Activity size={12} /> Live Market
                </Badge>
                <span className="text-zinc-500 text-[10px] font-mono uppercase tracking-tighter">
                  Binance Feed
                </span>
              </div>
              <CardTitle className="text-4xl font-bold tracking-tight">
                {symbol?.replace("USDT", "")}{" "}
                <span className="text-zinc-500 text-lg">/ USDT</span>
              </CardTitle>
            </CardHeader>

            <CardContent>
              <div
                className={`text-6xl font-mono font-bold transition-colors duration-500 ${priceColor}`}
              >
                $
                {parseFloat(data?.price || "0").toLocaleString(undefined, {
                  minimumFractionDigits: 2,
                })}
              </div>

              <div className="mt-6 flex gap-4">
                <div className="p-3 rounded-xl bg-zinc-900/50 border border-zinc-800 flex-1">
                  <p className="text-[10px] text-zinc-500 uppercase font-bold">
                    Estado
                  </p>
                  <span className="text-sm font-medium text-green-500 flex items-center gap-1">
                    Operativo{" "}
                    <span className="w-1.5 h-1.5 rounded-full bg-green-500 animate-pulse" />
                  </span>
                </div>
                <div className="p-3 rounded-xl bg-zinc-900/50 border border-zinc-800 flex-1">
                  <p className="text-[10px] text-zinc-500 uppercase font-bold">
                    API Latency
                  </p>
                  <p className="text-sm font-medium">24ms</p>
                </div>
              </div>
            </CardContent>
          </Card>

          {/* Sidebar de Stats o Lista Rápida */}
          <aside className="w-full md:w-72 space-y-4">
            <h3 className="text-xs font-bold text-zinc-500 uppercase tracking-widest px-1">
              Quick Watch
            </h3>
            <div className="grid gap-2">
              {["BTCUSDT", "ETHUSDT", "SOLUSDT"].map((s) => (
                <button
                  key={s}
                  onClick={() => (window.location.href = `/${s}`)}
                  className={`flex items-center justify-between p-3 rounded-lg border transition-all ${s === symbol ? "bg-blue-500/10 border-blue-500/50 text-blue-400" : "bg-zinc-900/30 border-zinc-800 hover:border-zinc-700"}`}
                >
                  <span className="font-bold text-xs">
                    {s.replace("USDT", "")}
                  </span>
                  {s === symbol ? (
                    <ArrowUpRight size={14} />
                  ) : (
                    <Globe size={14} className="opacity-20" />
                  )}
                </button>
              ))}
            </div>
          </aside>
        </div>
      </main>
    </div>
  );
}
