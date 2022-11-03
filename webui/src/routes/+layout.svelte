<script>
	import Header from '$lib/header/Header.svelte';
	import '../app.scss';
</script>

<Header />

<main>
	<slot />
</main>

<footer class="footer">
  <div class="level">
    <div class="level-left">
	  <p>visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to learn SvelteKit</p>
    </div>

    <div class="level-right">
      <p>Click on the logos to learn more</p>
    </div>
  </div>


</footer>

<style>
	main {
		padding: 1rem;
		width: 100%;
		max-width: 1024px;
		margin: 0 auto;
		box-sizing: border-box;
	}

	footer {
		justify-content: center;
		align-items: center;
		padding: 40px;
	}

	footer a {
		font-weight: bold;
	}

	@media (min-width: 480px) {
		footer {
			padding: 40px 0;
		}
	}
</style>
