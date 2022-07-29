// import { defineConfig } from "vite";
// // import { svelte } from "@sveltejs/vite-plugin-svelte";
import { sveltekit } from '@sveltejs/kit/vite';

// https://vitejs.dev/config/
// export default defineConfig({
//   plugins: [sveltekit()],
//   css: {
//     preprocessorOptions: {
//       scss: {
//         additionalData: '@use "src/variables.scss" as *;',
//       },
//     },
//   },
// });

/** @type {import('vite').UserConfig} */
const config = {
  plugins: [sveltekit()],
  build: {minify: false},
  css: {
	preprocessorOptions: {
	  scss: {
		additionalData: '@use "src/variables.scss" as *;'
	  }
	}
  },
  server: {fs: {strict: false}}
};

export default config;
