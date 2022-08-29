(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-05e15ad6"],{"1fc9":function(e,s,n){"use strict";n("4651")},"31e1":function(e,s,n){"use strict";n("f6d4")},4651:function(e,s,n){},6204:function(e,s,n){"use strict";n.r(s);var o=function(){var e=this,s=e.$createElement,n=e._self._c||s;return n("div",{staticClass:"login-container"},[n("el-form",{ref:"loginForm",staticClass:"login-form",attrs:{model:e.loginForm,rules:e.loginRules,"auto-complete":"on","label-position":"left"}},[n("div",{staticClass:"title-container"},[n("h3",{staticClass:"title"},[e._v(e._s(e.signupOrLogin))])]),n("el-form-item",{directives:[{name:"show",rawName:"v-show",value:!e.isLogin,expression:"!isLogin"}],attrs:{prop:"name"}},[n("span",{staticClass:"svg-container"},[n("svg-icon",{attrs:{"icon-class":"user"}})],1),n("el-input",{ref:"name",attrs:{placeholder:"Username",name:"name",type:"text",tabindex:"1","auto-complete":"on"},model:{value:e.loginForm.name,callback:function(s){e.$set(e.loginForm,"name",s)},expression:"loginForm.name"}})],1),n("el-form-item",{attrs:{prop:"email"}},[n("span",{staticClass:"svg-container"},[n("svg-icon",{attrs:{"icon-class":"email"}})],1),n("el-input",{ref:"email",attrs:{placeholder:"Enter your email",name:"email",type:"text",tabindex:"1","auto-complete":"on"},model:{value:e.loginForm.email,callback:function(s){e.$set(e.loginForm,"email",s)},expression:"loginForm.email"}})],1),n("el-form-item",{attrs:{prop:"password"}},[n("span",{staticClass:"svg-container"},[n("svg-icon",{attrs:{"icon-class":"password"}})],1),n("el-input",{key:e.passwordType,ref:"password",attrs:{type:e.passwordType,placeholder:"Password",name:"password",tabindex:"2","auto-complete":"on"},nativeOn:{keyup:function(s){return!s.type.indexOf("key")&&e._k(s.keyCode,"enter",13,s.key,"Enter")?null:e.handleLoginOrSignup(s)}},model:{value:e.loginForm.password,callback:function(s){e.$set(e.loginForm,"password",s)},expression:"loginForm.password"}}),n("span",{staticClass:"show-pwd",on:{click:e.showPwd}},[n("svg-icon",{attrs:{"icon-class":"password"===e.passwordType?"eye":"eye-open"}})],1)],1),n("el-form-item",{directives:[{name:"show",rawName:"v-show",value:!e.isLogin,expression:"!isLogin"}],attrs:{prop:"repeatPassword"}},[n("span",{staticClass:"svg-container"},[n("svg-icon",{attrs:{"icon-class":"password"}})],1),n("el-input",{key:e.passwordType,ref:"password",attrs:{type:e.passwordType,placeholder:"Repeat Your Password",name:"password",tabindex:"2","auto-complete":"on"},nativeOn:{keyup:function(s){return!s.type.indexOf("key")&&e._k(s.keyCode,"enter",13,s.key,"Enter")?null:e.handleLoginOrSignup(s)}},model:{value:e.loginForm.repeatPassword,callback:function(s){e.$set(e.loginForm,"repeatPassword",s)},expression:"loginForm.repeatPassword"}})],1),n("div",{staticClass:"btns"},[n("el-button",{staticStyle:{width:"60%","margin-bottom":"30px"},attrs:{loading:e.loading,type:"primary"},nativeOn:{click:function(s){return s.preventDefault(),e.handleLoginOrSignup(s)}}},[e._v(e._s(e.signupOrLogin))]),n("el-switch",{staticStyle:{width:"10%","margin-left":"5%"},model:{value:e.isLogin,callback:function(s){e.isLogin=s},expression:"isLogin"}}),n("span",{staticStyle:{color:"#fff"}},[e._v("shift to "+e._s(e.isLogin?"signup":"login"))])],1),n("div",{staticClass:"tips"},[e.isLogin?e._e():n("div",[e._v(" username: 1 to 30 characters")]),n("div",[e._v(" password: large than 6 digits")]),n("div",[e._v(" email: can be right, wrong or non-exist")])])],1)],1)},t=[],i=n("61f7"),a={name:"Login",data:function(){var e=function(e,s,n){Object(i["b"])(s)?n():n(new Error("Please enter the correct user name"))},s=function(e,s,n){s.length<6?n(new Error("The password can not be less than 6 digits")):n()};return{loginForm:{name:"",password:"",repeatPassword:"",email:""},loginRules:{name:[{trigger:"blur",validator:e}],password:[{required:!0,trigger:"blur",validator:s}]},loading:!1,passwordType:"password",redirect:void 0,isLogin:!0}},computed:{signupOrLogin:function(){return this.isLogin?"Login":"Signup"}},watch:{$route:{handler:function(e){this.redirect=e.query&&e.query.redirect},immediate:!0}},methods:{showPwd:function(){var e=this;"password"===this.passwordType?this.passwordType="":this.passwordType="password",this.$nextTick((function(){e.$refs.password.focus()}))},handleLoginOrSignup:function(){var e=this;this.$refs.loginForm.validate((function(s){if(!s)return console.log("error submit!!"),!1;e.loading=!0;var n=e.isLogin?"user/login":"user/signup";e.$store.dispatch(n,e.loginForm).then((function(){e.$router.push({path:e.redirect||"/"}),e.loading=!1})).catch((function(){e.loading=!1}))}))}}},r=a,l=(n("31e1"),n("1fc9"),n("2877")),c=Object(l["a"])(r,o,t,!1,null,"380e6ab4",null);s["default"]=c.exports},f6d4:function(e,s,n){}}]);