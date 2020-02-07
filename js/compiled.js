!function(t){var e={};function n(r){if(e[r])return e[r].exports;var o=e[r]={i:r,l:!1,exports:{}};return t[r].call(o.exports,o,o.exports,n),o.l=!0,o.exports}n.m=t,n.c=e,n.d=function(t,e,r){n.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:r})},n.r=function(t){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},n.t=function(t,e){if(1&e&&(t=n(t)),8&e)return t;if(4&e&&"object"==typeof t&&t&&t.__esModule)return t;var r=Object.create(null);if(n.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:t}),2&e&&"string"!=typeof t)for(var o in t)n.d(r,o,function(e){return t[e]}.bind(null,o));return r},n.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return n.d(e,"a",e),e},n.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},n.p="",n(n.s=5)}([function(t,e,n){},function(t,e,n){},function(t,e){t.exports=Vue},function(t,e){t.exports=Vuex},function(t,e){t.exports=VueRouter},function(t,e,n){t.exports=n(10)},function(t,e,n){},function(t,e){t.exports=jQuery},function(t,e,n){"use strict";var r=n(0);n.n(r).a},function(t,e,n){"use strict";var r=n(1);n.n(r).a},function(t,e,n){"use strict";n.r(e);n(6);var r=n(2),o=n.n(r),s=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"app"},[n("div",{staticClass:"header"},[n("h1",{on:{click:function(e){return t.gotoIndex()}}},[t._v("Go Vue Starter")]),t._v(" "),n("span",[t._v(t._s(t.username))]),t._v(" "),n("button",{attrs:{size:"mini"},on:{click:function(e){return t.logout()}}},[t._v("Log out")])]),t._v(" "),n("div",{staticClass:"content-area"},[n("router-view")],1)])};s._withStripped=!0;n(7);var u=n(3),i=n.n(u);o.a.use(i.a);var a=new i.a.Store({state:{},getters:{userID:function(t){return window.user.id},username:function(t){return window.user.username}}}),c=n(4),l=n.n(c),f=function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"index-page"},[t._v("\n\tVue app running\n\t"),n("br"),n("br"),t._v(" "),n("button",{on:{click:function(e){return t.testAjax()}}},[t._v("Test AJAX")])])};f._withStripped=!0;var p={methods:{testAjax:function(){this.$router.push({name:"test"})}}};n(8);function d(t,e,n,r,o,s,u,i){var a,c="function"==typeof t?t.options:t;if(e&&(c.render=e,c.staticRenderFns=n,c._compiled=!0),r&&(c.functional=!0),s&&(c._scopeId="data-v-"+s),u?(a=function(t){(t=t||this.$vnode&&this.$vnode.ssrContext||this.parent&&this.parent.$vnode&&this.parent.$vnode.ssrContext)||"undefined"==typeof __VUE_SSR_CONTEXT__||(t=__VUE_SSR_CONTEXT__),o&&o.call(this,t),t&&t._registeredComponents&&t._registeredComponents.add(u)},c._ssrRegister=a):o&&(a=i?function(){o.call(this,this.$root.$options.shadowRoot)}:o),a)if(c.functional){c._injectStyles=a;var l=c.render;c.render=function(t,e){return a.call(e),l(t,e)}}else{var f=c.beforeCreate;c.beforeCreate=f?[].concat(f,a):[a]}return{exports:t,options:c}}var v=d(p,f,[],!1,null,null,null);v.options.__file="src/pages/index.vue";var _=v.exports,m=function(){var t=this,e=t.$createElement;return(t._self._c||e)("div",{staticClass:"test-page"},[t.error?[0===t.error.readyState?[t._v("\n\t\t\tCould not connect\n\t\t")]:4===t.error.readyState?[t._v("\n\t\t\tError response code "+t._s(t.error.status)+"\n\t\t")]:[t._v("\n\t\t\tError in readyState "+t._s(t.error.readyState)+"\n\t\t")]]:[t._v("\n\t\t"+t._s(t.message)+"\n\t")]],2)};m._withStripped=!0;var h=d({data:function(){return{message:"",error:null}},beforeRouteEnter:function(t,e,n){$.get("/ajax/test").then((function(t){n((function(e){e.message=t.message}))})).fail((function(t){n((function(e){e.error=t}))}))},beforeRouteUpdate:function(t,e,n){$.get("/ajax/test").then((function(t){vm.message=t.message,vm.error=null,n()})).fail((function(t){vm.message="",vm.error=t,n()}))}},m,[],!1,null,null,null);h.options.__file="src/pages/test.vue";var g=h.exports,x=function(){var t=this.$createElement;return(this._self._c||t)("div",{staticClass:"not-found-page"},[this._v("\n\tNot found\n")])};x._withStripped=!0;var b=d({},x,[],!1,null,null,null);b.options.__file="src/pages/not-found.vue";var y=b.exports,S={store:a,router:new l.a({mode:"history",routes:[{name:"index",path:"/",component:_},{name:"test",path:"/test",component:g},{path:"*",component:y}]}),computed:{username:function(){return this.$store.getters.username}},methods:{gotoIndex:function(){"index"!==this.$route.name&&this.$router.push({name:"index"})},logout:function(){window.location.href="/logout"}}},w=(n(9),d(S,s,[],!1,null,null,null));w.options.__file="src/app.vue";var C=w.exports;new o.a(C).$mount("#app")}]);