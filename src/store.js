import { createStore } from 'vuex';
import ajax from './ajax.js';

export default createStore({
	state() {
		return {
			user: null,
		};
	},
	getters: {
		authenticated(state) {
			return !!state.user;
		},
		username(state) {
			return state.user ? state.user.username : '';
		},
	},
	mutations: {
		setUser(state, user) {
			state.user = user;
		},
	},
	actions: {
		checkLogin({commit}) {
			return ajax.get('/ajax/fetchLogin').then((response) => {
				commit('setUser', response.data);
			});
		},
		logIn({getters, dispatch, commit}, postParams) {
			if (getters.authenticated) {
				return dispatch('logOut').then(() => dispatch('logIn', postParams));
			}
			return ajax.post('/ajax/login', postParams).then((response) => {
				commit('setUser', response.data);
			});
		},
		logOut({getters, commit}) {
			if (!getters.authenticated) {
				return Promise.resolve();
			}
			return ajax.post('/ajax/logout').then(() => {
				commit('setUser', null);
			});
		},
	},
});
