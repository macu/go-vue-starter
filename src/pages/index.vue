<template>
<div class="index-page">
	Vue app running
	<br/><br/>
	<button @click="testAjax()">Test AJAX</button>
</div>
</template>

<script>
export default {
	methods: {
		testAjax() {
			this.$router.push({name: 'test'});
		},
	},
};
</script>

<style lang="scss">
.index-page {
}
</style>
