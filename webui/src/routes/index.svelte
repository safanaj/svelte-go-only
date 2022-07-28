<script context="module">
	export const prerender = true;
</script>

<script>
  import Counter from '$lib/Counter.svelte';
  import viteLogo from '$lib/assets/vite.svg'
  import svelteLogo from '$lib/assets/svelte.svg'
  import ginLogo from '$lib/assets/gin-gonic.svg'
  import grpcLogo from '$lib/assets/grpc.svg'
  import Category from '$lib/Category.svelte'
  import Streaming from '$lib/Streaming.svelte'

  import { Container, Col, Row, Input, InputGroup } from 'sveltestrap';
</script>

<svelte:head>
	<!-- <title>Home</title> -->
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<div class="main-content">
  <Container>
    <Row>
      <div class="header-title">
        <Row>
          <Col>
          <a href="https://vitejs.dev" target="_blank">
            <img src={viteLogo} class="logo" alt="Vite Logo" />
          </a>
        </Col>
        <Col>
        <a href="https://svelte.dev" target="_blank">
          <img src={svelteLogo} class="logo svelte" alt="Svelte Logo" />
        </a>
        </Col>
        <Col>
        <a href="https://gin-gonic.com" target="_blank">
          <img src={ginLogo} class="logo gin-gonic" alt="Gin Gonic Logo" />
        </a>
        </Col>
        <Col>
        <a href="https://grpc.io" target="_blank">
          <img src={grpcLogo} class="logo grpc" alt="gRPC Logo" />
        </a>
        </Col>
        </Row>

      </div>

    <h1>Vite + Svelte by Go Gin + GRPC-Web</h1>
  </Row>

  <Row>
    <div id="streaming" class="box">
      <Streaming />
    </div>
  </Row>

  <Row>
    <div id="category" class="card">
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

</Container>
</div>

<style>
  #streaming { justify-content: center; text-align: center; witdth: 100%; }
  #category { justify-content: center; text-align: center; witdth: 100%; }

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
