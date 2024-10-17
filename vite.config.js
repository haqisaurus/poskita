import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import path from "path";

export default defineConfig({
    plugins: [
        react(),
    ],
    resolve: {
        alias: {
            "@": path.resolve(__dirname, "src/react"),
        },
    },
    build: {
        assetsInlineLimit: 0,
        emptyOutDir: true,
        rollupOptions: {
            output: {
                manifest: true,
            },
            
        },
        outDir: "bin"
    },
    
});
