<script context="module">
	export const prerender = true;
</script>

<script>
  import Counter from '$lib/Counter.svelte';
  import viteLogo from '$lib/assets/vite.svg'
  import svelteLogo from '$lib/assets/svelte.svg'
  import ginLogo from '$lib/assets/gin-gonic.svg'
  import grpcLogo from '$lib/assets/grpc.svg'
  // import Counter from './lib/Counter.svelte'
  import Category from '$lib/Category.svelte'

  import { Col, Row, Input, InputGroup } from 'sveltestrap';

</script>

<svelte:head>
	<!-- <title>Home</title> -->
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<!-- <section> -->
<!-- 	<h1> -->
<!-- 		<span class="welcome"> -->
<!-- 			<picture> -->
<!-- 				<source srcset="svelte-welcome.webp" type="image/webp" /> -->
<!-- 				<img src="svelte-welcome.png" alt="Welcome" /> -->
<!-- 			</picture> -->
<!-- 		</span> -->

<!-- 		to your new<br />SvelteKit app -->
<!-- 	</h1> -->

<!-- 	<h2> -->
<!-- 		try editing <strong>src/routes/index.svelte</strong> -->
<!-- 	</h2> -->

<!-- 	<Counter /> -->
<!-- </section> -->

<!-- <style> -->
<!-- 	section { -->
<!-- 		display: flex; -->
<!-- 		flex-direction: column; -->
<!-- 		justify-content: center; -->
<!-- 		align-items: center; -->
<!-- 		flex: 1; -->
<!-- 	} -->

<!-- 	h1 { -->
<!-- 		width: 100%; -->
<!-- 	} -->

<!-- 	.welcome { -->
<!-- 		display: block; -->
<!-- 		position: relative; -->
<!-- 		width: 100%; -->
<!-- 		height: 0; -->
<!-- 		padding: 0 0 calc(100% * 495 / 2048) 0; -->
<!-- 	} -->

<!-- 	.welcome img { -->
<!-- 		position: absolute; -->
<!-- 		width: 100%; -->
<!-- 		height: 100%; -->
<!-- 		top: 0; -->
<!-- 		display: block; -->
<!-- 	} -->
<!-- </style> -->

<section>
  <Row>
    <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} class="logo" alt="Vite Logo" />
        </a>
        <a href="https://svelte.dev" target="_blank">
          <img src={svelteLogo} class="logo svelte" alt="Svelte Logo" />
        </a>
        <a href="https://gin-gonic.com" target="_blank">
          <img src={ginLogo} class="logo gin-gonic" alt="Gin Gonic Logo" />
        </a>
        <a href="https://grpc.io" target="_blank">
          <img src={grpcLogo} class="logo grpc" alt="gRPC Logo" />
        </a>
    </div>

    <h1>Vite + Svelte by Go Gin + GRPC-Web</h1>
  </Row>

    <!-- <Row> -->

    <!--   <Col xs="4"> -->
    <!--     <div class="card"> -->
    <!--       <Counter /> -->
    <!--     </div> -->
    <!--   </Col> -->

    <!--   <Col pull="right"> -->
    <!--     <div class="card"> -->
    <!--       <Counter /> -->
    <!--     </div> -->
    <!--   </Col> -->

    <!-- </Row> -->

    <Row>
      <div class="card">
        <Category />
      </div>
    </Row>


    <Row>
      <p>
        Check out <a href="https://github.com/sveltejs/kit#readme" target="_blank">SvelteKit</a>, the official Svelte app framework powered by Vite!
      </p>

      <p class="read-the-docs">
        Click on the Vite and Svelte logos to learn more
      </p>
    </Row>

</section>

<style>
  .logo {
      height: 6em;
      padding: 1.5em;
      will-change: filter;
  }
  .logo:hover {
      filter: drop-shadow(0 0 2em #646cffaa);
  }
  .logo.svelte:hover {
      filter: drop-shadow(0 0 2em #ff3e00aa);
  }
  .read-the-docs {
      color: #888;
  }
</style>
