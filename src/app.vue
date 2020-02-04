<template>
<div class="app">
	<div class="header">
		<h1>Go Vue Starter</h1>
		<button @click="logout()" size="mini">Log out</button>
	</div>
	<div class="content-area">
		Vue app running
		<br/><br/>
		Username: {{username}}
		<br/><br/>
		<button @click="testAjax()">Test AJAX</button>
	</div>
</div>
</template>

<script>
import $ from 'jquery';
import store from './store.js';

export default {
	store,
	computed: {
		username() {
			return this.$store.getters.username;
		},
	},
	methods: {
		testAjax() {
			$.get('/ajax/test').then(response => {
				alert(response.message);
			}).fail(jqXHR => {
				alert(jqXHR.status);
			});
		},
		logout() {
			window.location.href = '/logout';
		},
	},
};
</script>

<style lang="scss">
.app {
	height: 100%;
	>.header {
		display: flex;
		background-color: lightblue;
		padding: 20px 1cm;
		align-items: center;
		>h1 {
			flex: 1;
			margin: 0;
		}
	}
	>.content-area {
		padding: 1cm;
	}
}
</style>
