<script>
	import Header from './Header.svelte';
	import Footer from './Footer.svelte';
	import '../app.css';
	import './styles.css'
</script>

<div class="app">
	<Header />

	<main>
		<slot />
	</main>
	<Footer />
</div>