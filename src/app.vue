<template>
<div class="app">
	<div class="header">
		<h1 @click="gotoIndex()">Go Vue Starter</h1>
		<template v-if="authenticated">
			<span>{{username}}</span>
			<button @click="logOut()">Log out</button>
		</template>
		<span v-else-if="loadingLogin">Loading login...</span>
		<button v-else @click="logIn()">Log in</button>
	</div>
	<div class="content-area">
		<router-view></router-view>
	</div>
</div>
</template>

<script>
export default {
	data() {
		return {
			loadingLogin: true,
		};
	},
	computed: {
		authenticated() {
			return this.$store.getters.authenticated;
		},
		username() {
			return this.$store.getters.username;
		},
	},
	mounted() {
		this.$store.dispatch('checkLogin').finally(() => {
			this.loadingLogin = false;
		});
	},
	methods: {
		gotoIndex() {
			if (this.$route.name !== 'index') {
				this.$router.push({name: 'index'});
			}
		},
		logIn() {
			let username = window.prompt('Username', '');
			if (!(username || []).trim()) {
				return;
			}
			let password = window.prompt('Password', '');
			if (!(password || []).trim()) {
				return;
			}
			this.loadingLogin = true;
			this.$store.dispatch('logIn', {username, password}).catch(err => {
				if (err && err.response && err.response.status === 403) {
					window.alert('Invalid username or password');
				} else {
					window.alert('Error logging in');
				}
			}).finally(() => {
				this.loadingLogin = false;
			});
		},
		logOut() {
			if (window.confirm('Log out?')) {
				this.$store.dispatch('logOut').catch(err => {
					window.alert('Error logging out');
				});
			}
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
