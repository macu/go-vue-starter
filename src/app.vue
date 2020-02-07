<template>
<div class="app">
	<div class="header">
		<h1 @click="gotoIndex()">Go Vue Starter</h1>
		<span>{{username}}</span>
		<button @click="logout()" size="mini">Log out</button>
	</div>
	<div class="content-area">
		<router-view></router-view>
	</div>
</div>
</template>

<script>
import $ from 'jquery';
import store from './store.js';
import router from './router.js';

export default {
	store,
	router,
	computed: {
		username() {
			return this.$store.getters.username;
		},
	},
	methods: {
		gotoIndex() {
			if (this.$route.name !== 'index') {
				this.$router.push({name: 'index'});
			}
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
			cursor: pointer;
		}
		>*:not(:last-child) {
			margin-right: 20px;
		}
	}
	>.content-area {
		padding: 1cm;
	}
}
</style>
