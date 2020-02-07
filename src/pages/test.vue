<template>
<div class="test-page">
	<template v-if="error">
		<template v-if="error.readyState === 0">
			Could not connect
		</template>
		<template v-else-if="error.readyState === 4">
			Error response code {{error.status}}
		</template>
		<template v-else>
			Error in readyState {{error.readyState}}
		</template>
	</template>
	<template v-else>
		{{message}}
	</template>
</div>
</template>

<script>
export default {
	data() {
		return {
			message: '',
			error: null,
		};
	},
	beforeRouteEnter (to, from, next) {
		$.get('/ajax/test').then(response => {
			next(vm => {
				vm.message = response.message;
			});
		}).fail(jqXHR => {
			next(vm => {
				vm.error = jqXHR;
			});
		});
  },
	beforeRouteUpdate(to, from, next) {
		$.get('/ajax/test').then(response => {
			vm.message = response.message;
			vm.error = null;
			next();
		}).fail(jqXHR => {
			vm.message = '';
			vm.error = jqXHR;
			next();
		});
	},
};
</script>
