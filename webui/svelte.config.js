import preprocess from "svelte-preprocess";
import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
  kit: {
    adapter: adapter(),
    outDir: 'dist',
    // prerender: {
    //   crawl: false,
    //   default: true
    // },

    // // Override http methods in the Todo forms
    // methodOverride: {
    //   allowed: ['PATCH', 'DELETE']
    // }
  },
  preprocess: [
    preprocess({
      scss: {
        prependData: '@use "src/variables.scss" as *;',
      },
    }),
  ],
};

export default config;
