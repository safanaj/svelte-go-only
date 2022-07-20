<script>
  import viteLogo from './assets/vite.svg'
  import svelteLogo from './assets/svelte.svg'
  import ginLogo from './assets/gin-gonic.svg'
  import grpcLogo from './assets/grpc.svg'
  // import Counter from './lib/Counter.svelte'
  import Category from './lib/Category.svelte'

  import { Container, Col, Row, Input, InputGroup } from 'sveltestrap';
</script>

<main>
  <Container fluid>
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
  </Container>
</main>

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
