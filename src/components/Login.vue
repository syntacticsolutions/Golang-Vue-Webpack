<template>
<div class="text-center" style="padding:50px 0">	
	<!-- Main Form -->
	<div class="login-form-1">
        <div class="logo text-left">
            <h5 v-if="action === 'login'">Project Map Login</h5>
            <h5 v-if="action === 'register'">Registration</h5>
            <h5 v-if="action === 'forgot'">Forgot Password</h5>
        </div>
		<div id="login-form" class="text-left">
			<div class="login-form-main-message"></div>
			<div class="main-login-form">
				<div class="login-group">
                    <q-input v-if="action === 'register'" v-model="first_name" float-label="First Name" placeholder="First Name" />
                    <q-input v-if="action === 'register'" v-model="last_name" float-label="Last Name" placeholder="Last Name" />
                    <q-input v-model="email" float-label="Email" placeholder="Email" />
                    <q-input v-if="action !== 'forgot'" type="password" v-model="password" float-label="Password" placeholder="Password" />
                    <q-input v-if="action === 'register'" type="password" v-model="confirm_password" float-label="Confirm Password" placeholder="Confirm Password" />
					<div class="form-group login-group-checkbox" v-if="action !== 'forgot'">
						<input type="checkbox" id="lg_remember" name="lg_remember">
						<label for="lg_remember">remember</label>
					</div>
				</div>
				<button @keyup.enter="submit()" @click="submit()" class="login-button"><i class="fa fa-chevron-right"></i></button>
			</div>
			<div class="etc-login-form" v-if="action === 'login'">
				<p>forgot your password? <a @click="setForm('forgot')">click here</a></p>
				<p>new user? <a @click="setForm('register')">create new account</a></p>
			</div>
            <div class="etc-login-form" v-if="action === 'register'">
				<p>already have an account? <a @click="setForm('login')">sign in</a></p>
			</div>
            <div class="etc-login-form" v-if="action === 'forgot'">
				<p>new user? <a @click="setForm('register')">create new account</a></p>
                <p>login? <a @click="setForm('login')">click here</a></p>
			</div>
		</div>
	</div>
	<!-- end:Main Form -->
</div>
</template>

<script>
import { QInput } from 'quasar'
import axios from 'axios';

const _ = {
    cloneDeep: require('lodash/cloneDeep'),
    omit: require('lodash/omit')
}

const endpoints = {
    login: '/login',
    register: '/register',
    forgot: '/forgot_password'
}

export default {
    components: {
        QInput
    },
    data () {
        return {
            action:'login',
            email: null,
            password: null,
            confirm_password: null,
            first_name: null,
            last_name: null
        }
    },
    mounted() {
        window.login = this;
    },
    methods: {
        setForm(val) {
            this.action = val
        },
        submit() {
            const apiString = endpoints[this.action];
            const payload = this.getPayload(this.action);

            axios.post(config.host + apiString, payload)
            .then(res =>{
                switch(this.action){
                    case 'login': this.login(res);
                    case 'register': this.register(res);
                    case 'forgot': this.sendForgotEmail(res);
                }
            }, err => {
                console.log(err);
            })
        },
        getPayload(action) {
            switch (action){

                case 'login':
                return {
                    email: this.email,
                    password: this.password
                }

                case 'register':
                return _.omit(_.cloneDeep(this._data), 'action')

                case 'forgot':
                return {
                    email: this.email
                }
            }
        },
        login(res) {
            window.axios.defaults.headers.common['Authorization'] = 'Bearer ' + res.data.token;
            this.$root.user.id = res.data.id;
            this.$root.user.name = res.data.name

            this.$root.$router.replace('/admin')
            // window.axios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
        },
        register(res) {
            window.axios.defaults.headers.common['Authorization']  = 'Bearer ' + res.data.token;

        },
        sendForgotEmail() {

        }

    }
}
</script>

<style scoped lang="scss">
/*------------------------------------------------------------------
[Master Stylesheet]

Project    	: Aether
Version		: 1.0
Last change	: 2015/03/27
-------------------------------------------------------------------*/
/*------------------------------------------------------------------
[Table of contents]

1. General Structure
2. Anchor Link
3. Text Outside the Box
4. Main Form
5. Login Button
6. Form Invalid
7. Form - Main Message
8. Custom Checkbox & Radio
9. Misc
-------------------------------------------------------------------*/
/*=== 1. General Structure ===*/
html,
body {
  background: #efefef;
  padding: 10px;
  font-family: 'Varela Round';
}
/*=== 2. Anchor Link ===*/
a {
  color: #aaaaaa;
  transition: all ease-in-out 200ms;
}
a:hover {
  color: #333333;
  text-decoration: none;
}
/*=== 3. Text Outside the Box ===*/
.etc-login-form {
  color: #919191;
  padding: 10px 20px;
}
.etc-login-form p {
  margin-bottom: 5px;
}
/*=== 4. Main Form ===*/
.login-form-1 {
    max-width: 300px;
    border-radius: 5px;
    display: inline-block;
    margin-right: 25px;
    width: 285px;
}
.main-login-form {
  position: relative;
}
.login-form-1 .form-control {
  border: 0;
  box-shadow: 0 0 0;
  border-radius: 0;
  background: transparent;
  color: #555555;
  padding: 7px 0;
  font-weight: bold;
  height:auto;
}
.login-form-1 .form-control::-webkit-input-placeholder {
  color: #999999;
}
.login-form-1 .form-control:-moz-placeholder,
.login-form-1 .form-control::-moz-placeholder,
.login-form-1 .form-control:-ms-input-placeholder {
  color: #999999;
}
.login-form-1 .form-group {
  margin-bottom: 0;
  border-bottom: 2px solid #efefef;
  padding-right: 20px;
  position: relative;
}
.login-form-1 .form-group:last-child {
  border-bottom: 0;
}
.login-group {
  background: #ffffff;
  color: #999999;
  border-radius: 8px;
  padding: 10px 25px;
}
.login-group-checkbox {
  padding: 5px 0;
}
/*=== 5. Login Button ===*/
.login-form-1 .login-button {
  position: absolute;
  right: -25px;
  top: 50%;
  background: #ffffff;
  color: #999999;
  padding: 11px 0;
  width: 50px;
  height: 50px;
  margin-top: -25px;
  border: 5px solid #efefef;
  border-radius: 50%;
  transition: all ease-in-out 500ms;
}
.login-form-1 .login-button:hover {
  color: #555555;
  transform: rotate(450deg);
}
.login-form-1 .login-button.clicked {
  color: #555555;
}
.login-form-1 .login-button.clicked:hover {
  transform: none;
}
.login-form-1 .login-button.clicked.success {
  color: #2ecc71;
}
.login-form-1 .login-button.clicked.error {
  color: #e74c3c;
}
/*=== 6. Form Invalid ===*/
label.form-invalid {
  position: absolute;
  top: 0;
  right: 0;
  z-index: 5;
  display: block;
  margin-top: -25px;
  padding: 7px 9px;
  background: #777777;
  color: #ffffff;
  border-radius: 5px;
  font-weight: bold;
  font-size: 11px;
}
label.form-invalid:after {
  top: 100%;
  right: 10px;
  border: solid transparent;
  content: " ";
  height: 0;
  width: 0;
  position: absolute;
  pointer-events: none;
  border-color: transparent;
  border-top-color: #777777;
  border-width: 6px;
}
/*=== 7. Form - Main Message ===*/
.login-form-main-message {
  background: #ffffff;
  color: #999999;
  border-left: 3px solid transparent;
  border-radius: 3px;
  margin-bottom: 8px;
  font-weight: bold;
  height: 0;
  padding: 0 20px 0 17px;
  opacity: 0;
  transition: all ease-in-out 200ms;
}
.login-form-main-message.show {
  height: auto;
  opacity: 1;
  padding: 10px 20px 10px 17px;
}
.login-form-main-message.success {
  border-left-color: #2ecc71;
}
.login-form-main-message.error {
  border-left-color: #e74c3c;
}
/*=== 8. Custom Checkbox & Radio ===*/
/* Base for label styling */
[type="checkbox"]:not(:checked),
[type="checkbox"]:checked,
[type="radio"]:not(:checked),
[type="radio"]:checked {
  position: absolute;
  left: -9999px;
}
[type="checkbox"]:not(:checked) + label,
[type="checkbox"]:checked + label,
[type="radio"]:not(:checked) + label,
[type="radio"]:checked + label {
  position: relative;
  padding-left: 25px;
  padding-top: 1px;
  cursor: pointer;
}
/* checkbox aspect */
[type="checkbox"]:not(:checked) + label:before,
[type="checkbox"]:checked + label:before,
[type="radio"]:not(:checked) + label:before,
[type="radio"]:checked + label:before {
  content: '';
  position: absolute;
  left: 0;
  top: 2px;
  width: 17px;
  height: 17px;
  border: 0px solid #aaa;
  background: #f0f0f0;
  border-radius: 3px;
  box-shadow: inset 0 1px 3px rgba(0, 0, 0, 0.3);
}
/* checked mark aspect */
[type="checkbox"]:not(:checked) + label:after,
[type="checkbox"]:checked + label:after,
[type="radio"]:not(:checked) + label:after,
[type="radio"]:checked + label:after {
  position: absolute;
  color: #555555;
  transition: all .2s;
}
/* checked mark aspect changes */
[type="checkbox"]:not(:checked) + label:after,
[type="radio"]:not(:checked) + label:after {
  opacity: 0;
  transform: scale(0);
}
[type="checkbox"]:checked + label:after,
[type="radio"]:checked + label:after {
  opacity: 1;
  transform: scale(1);
}
/* disabled checkbox */
[type="checkbox"]:disabled:not(:checked) + label:before,
[type="checkbox"]:disabled:checked + label:before,
[type="radio"]:disabled:not(:checked) + label:before,
[type="radio"]:disabled:checked + label:before {
  box-shadow: none;
  border-color: #8c8c8c;
  background-color: #878787;
}
[type="checkbox"]:disabled:checked + label:after,
[type="radio"]:disabled:checked + label:after {
  color: #498C8C;
}
[type="checkbox"]:disabled + label,
[type="radio"]:disabled + label {
  color: #498C8C;
}
/* accessibility */
[type="checkbox"]:checked:focus + label:before,
[type="checkbox"]:not(:checked):focus + label:before,
[type="checkbox"]:checked:focus + label:before,
[type="checkbox"]:not(:checked):focus + label:before {
  border: 1px dotted #f6f6f6;
}
/* hover style just for information */
label:hover:before {
  border: 1px solid #498C8C !important;
}
/*=== Customization ===*/
/* radio aspect */
[type="checkbox"]:not(:checked) + label:before,
[type="checkbox"]:checked + label:before {
  border-radius: 3px;
}
[type="radio"]:not(:checked) + label:before,
[type="radio"]:checked + label:before {
  border-radius: 35px;
}
/* selected mark aspect */
[type="checkbox"]:not(:checked) + label:after,
[type="checkbox"]:checked + label:after {
  content: '✔';
  top: 0;
  left: 2px;
  font-size: 14px;
}
[type="radio"]:not(:checked) + label:after,
[type="radio"]:checked + label:after {
  content: '\2022';
  top: 0;
  left: 3px;
  font-size: 30px;
  line-height: 25px;
}
/*=== 9. Misc ===*/
.logo {
  padding: 15px 20px;
  font-size: 25px;
  color: #498C8C;
  font-weight: 700;
  position:relative;
  white-space:nowrap;
}

</style>