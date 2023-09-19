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
import ajax from '../ajax.js';

export default {
	data() {
		return {
			message: '',
			error: null,
		};
	},
	beforeRouteEnter (to, from, next) {
		ajax.get('/ajax/test').then(response => {
			next(vm => {
				vm.message = response.data.message;
			});
		}).catch(err => {
			next(vm => {
				vm.error = err.toJSON();
			});
		});
	},
	beforeRouteUpdate(to, from, next) {
		ajax.get('/ajax/test').then(response => {
			vm.message = response.data.message;
			vm.error = null;
			next();
		}).catch(err => {
			vm.message = '';
			vm.error = err.toJSON();
			next();
		});
	},
};
</script>
