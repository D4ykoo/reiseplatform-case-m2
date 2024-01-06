import { fileURLToPath, URL } from "node:url";

import { defineConfig, loadEnv } from "vite";
import vue from "@vitejs/plugin-vue";
import ImportMetaEnvPlugin from "@import-meta-env/unplugin";

// export default defineConfig(({ command, mode }) => {
//   // Load env file based on `mode` in the current working directory.
//   // Set the third parameter to '' to load all env regardless of the `VITE_` prefix.
//   const env = loadEnv(mode, process.cwd(), '')
//   return {
//     // vite config
//     define: {
//       __APP_ENV__: JSON.stringify(env.APP_ENV),
//     },
//     plugins: [
//       vue(),
//       ImportMetaEnvPlugin.vite({
//         example: ".env.example",
//         // "env": "...",
//         // "transformMode": "..."
//       }),
//     ],
//     resolve: {
//       alias: {
//         "@": fileURLToPath(new URL("./src", import.meta.url)),
//       },
//     },
//   }
// })

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    ImportMetaEnvPlugin.vite({
      example: ".env.example.public",
      env: ".env",
      // "transformMode": "..."
    }),
  ],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
});
