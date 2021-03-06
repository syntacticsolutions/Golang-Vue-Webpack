<template>
<q-list>
  <q-collapsible icon="explore" label="Project Info">
    <div>
      <q-input v-model="project.title" float-label="Project Title"/>
    </div>
  </q-collapsible>
  <q-collapsible icon="fa-calendar" label="Project Dates">
    <div>
      <q-datetime-range
	  	@change="changeDateRange"
        type="date"
        v-model="range"
		format="MM/DD/YYYY"
		float-label="Start Date"
        />
    </div>
  </q-collapsible>
    <q-collapsible icon="fa-twitter" label="Project Manager">
    <div>
        <q-select
          v-model="project.pm_id"
          float-label="Project Manager"
          :options="pmOptions"
          @change="changePM"
        />
		<q-input v-model="project.project_manager_email" float-label="Email" :disable="true" />
		<q-input v-model="project.project_manager_phone" float-label="Office Phone" :disable="true"/>
		<q-input v-model="project.project_manager_cell" float-lavel="Cell" :disable="true"/>
    </div>
  </q-collapsible>
    <q-collapsible icon="fa-github" label="Contractor">
    <div>
      <q-select
          v-model="project.contractor_id"
          float-label="Contractor"
          :options="contractorOptions"
          @change="changeContractor"
        />
		<q-input v-model="project.contractor_email" float-label="Email" :disable="true" />
		<q-input v-model="project.contractor_phone" float-label="Office Phone" :disable="true"/>
		<q-input v-model="project.contractor_cell" float-lavel="Cell" :disable="true"/>
    </div>
  </q-collapsible>
  <q-collapsible icon="location_on" label="Marker(s)">
    <div>
		<q-chip v-for="(marker, id) in project.markers"
			closable
			:key="id"
			:style="'background:' +  project.color + ';color:#f1f1f1;'"
			@close="confirmDeleteMarker(marker)">
			{{ marker.type.split('_').map(res=>res.capitalize()).join(' ') }}
		</q-chip>
    </div>
  </q-collapsible>
  <div class="row accordion-actions">
	  <div><q-icon name="delete" size="29px" color="danger"></q-icon><p class="text-danger">Delete</p></div>
	  <div><q-icon name="fa-arrow-left" size="29px" color="warn"></q-icon><p class="text-warn">Cancel</p></div>
	  <div><q-icon name="save" size="29px" color="primary"></q-icon><p class="text-primary">Save</p></div>

  </div>
</q-list>
</template>

<script>

import { 
	QList,
	QCollapsible,
	QInput,
	QDatetimeRange,
	QSelect,
	QChip,
	QIcon } from 'quasar';

const _ = {
	each: require('lodash/each'),
	isEmpty: require('lodash/isEmpty')
}

export default {
    props: ['project'],
    components: {
        QList,
        QCollapsible,
        QInput,
        QDatetimeRange,
        QSelect,
		QChip,
		QIcon
    },
	data(){
		return {
			project_managers: {},
			pmOptions: [],
			contractorOptions: [],
			contractors: {}
		}
	},
    created(){

    },
    computed: {
      // TODO
        range() {
            return {
                from: this.project.start_date.String,
                to: this.project.end_date.String
            }
        }
    },
	methods: {
		changePM(id){
			this.project.project_manager = this.project_managers[id].project_manager;
			this.project.pm_id = id;
			this.project.project_manager_cell = this.project_managers[id].project_manager_cell
			this.project.project_manager_phone = this.project_managers[id].project_manager_phone
			this.project.project_manager_email = this.project_managers[id].project_manager_email
		},
		changeContractor(id){
			this.project.contractor = this.contractors[id].contractor;
			this.project.contractor_id = id;
			this.project.contractor_cell = this.contractors[id].contractor_cell
			this.project.contractor_phone = this.contractors[id].contractor_phone
			this.project.contractor_email = this.contractors[id].contractor_email
			this.project.color = this.contractors[id].color
		},
		changeDateRange(event){
			if(event.from){
				this.project.start_date.String = event.from.split('.')[0].split('T').join(' ')
			}
			if(event.to){
				this.project.end_date.String = event.to.split('.')[0].split('T').join(' ')
			}
		},
		confirmDeleteMarker(marker){

		}
	},
	mounted(){
		window.accord = this;
		axios.get(config.host + '/api/project_managers')
		.then(res => {
			var self = this;
			_.each(res.data.project_managers, (pm)=>{
				self.project_managers[pm.id] = pm;
				self.pmOptions.push({label: pm.project_manager, value: pm.id});
			})
		})

		axios.get(config.host + '/api/contractors')
		.then(res => {
			var self = this;
			_.each(res.data.contractors, (c)=>{
				self.contractors[c.id] = c;
				self.contractorOptions.push({label: c.contractor, value: c.id});
			})
		})

		$('.q-datetime-range-right')[0].children[0].children[0].innerHTML = 'End Date'
	}

}

String.prototype.capitalize = function(){
	let firstLetter = this.charAt(0).toUpperCase()
	return firstLetter + this.slice(1)
}

</script>
<style lang="scss">
.accordion-actions{
	padding-top:20px;
	border-top:1px solid #e0e0e0;
	div {
		min-width: 100px;
		text-align:center;
	}
}

.fa-github:before {
    content: url("data:image/svg+xml,%3Csvg%20xmlns%3D%22http%3A//www.w3.org/2000/svg%22%20xmlns%3Axlink%3D%22http%3A//www.w3.org/1999/xlink%22%20version%3D%221.1%22%20id%3D%22Layer_1%22%20x%3D%220px%22%20y%3D%220px%22%20viewBox%3D%220%200%20503.607%20503.607%22%20style%3D%22enable-background%3Anew%200%200%20503.607%20503.607%3B%22%20xml%3Aspace%3D%22preserve%22%20width%3D%2230px%22%20height%3D%2230px%22%3E%3Cg%3E%3Cg%3E%0A%09%3Cg%3E%0A%09%09%3Cg%3E%0A%09%09%09%3Cpath%20d%3D%22M478.67%2C417.993c-50.352-26.498-88.87-41.212-123.081-47.02l3.727-5.12c1.46-2.014%2C1.956-4.558%2C1.36-6.967l-7.798-31.165%20%20%20%20%20c20.883-29.444%2C33.221-66.224%2C33.221-101.099v-33.574h8.393h16.787c4.642%2C0%2C8.393-3.76%2C8.393-8.393%20%20%20%20%20c0-4.642-3.752-8.393-8.393-8.393h-8.393v-8.393c0-23.384-5.506-45.409-15.377-64.999l-75.365%2C14.487%20%20%20%20%20c-0.529%2C0.101-1.058%2C0.151-1.586%2C0.151c-2.199%2C0-4.339-0.873-5.926-2.451c-1.964-1.972-2.837-4.784-2.317-7.521l13.765-72.268%20%20%20%20%20c-7.109-3.198-14.537-5.968-22.31-8.175v-1.779C293.77%2C11.356%2C282.414%2C0%2C268.456%2C0h-33.297%20%20%20%20%20c-13.967%2C0-25.323%2C11.356-25.323%2C25.315v1.779c-7.764%2C2.208-15.192%2C4.977-22.301%2C8.175l13.765%2C72.268%20%20%20%20%20c0.52%2C2.736-0.353%2C5.548-2.317%2C7.521c-1.586%2C1.578-3.727%2C2.451-5.934%2C2.451c-0.529%2C0-1.058-0.05-1.586-0.151l-75.356-14.487%20%20%20%20%20c-9.871%2C19.59-15.385%2C41.615-15.385%2C64.999v8.393h-8.393c-4.633%2C0-8.393%2C3.752-8.393%2C8.393c0%2C4.633%2C3.76%2C8.393%2C8.393%2C8.393%20%20%20%20%20h16.787h8.393v33.574c0%2C34.875%2C12.347%2C71.655%2C33.221%2C101.099l-7.789%2C31.165c-0.604%2C2.409-0.101%2C4.952%2C1.36%2C6.967l3.727%2C5.12%20%20%20%20%20c-34.212%2C5.808-72.729%2C20.522-123.081%2C47.02C9.56%2C426.085%2C0%2C441.923%2C0%2C459.323v35.89c0%2C4.633%2C3.76%2C8.393%2C8.393%2C8.393h201.443%20%20%20%20%20h0.008h3.139c0.008%2C0%2C0.008-0.008%2C0.017-0.008l77.253-0.067c0.134%2C0.008%2C0.243%2C0.076%2C0.378%2C0.076h204.582%20%20%20%20%20c4.642%2C0%2C8.393-3.76%2C8.393-8.393v-35.89C503.607%2C441.923%2C494.055%2C426.085%2C478.67%2C417.993z%20M134.295%2C226.623v-33.574h235.016%20%20%20%20%20v33.574c0%2C37.796-13.371%2C71.059-33.146%2C96.004c-0.294%2C0.285-0.68%2C0.462-0.932%2C0.789c-2.887%2C3.693-5.875%2C7.176-8.947%2C10.45%20%20%20%20%20c-0.017%2C0.017-0.034%2C0.034-0.05%2C0.05c-2.04%2C2.157-4.163%2C4.079-6.278%2C6.052c-20.514%2C18.499-44.88%2C29.344-68.155%2C29.344%20%20%20%20%20c-23.267%2C0-47.633-10.844-68.155-29.344c-2.107-1.972-4.23-3.895-6.27-6.052c-0.017-0.017-0.034-0.034-0.05-0.05%20%20%20%20%20c-3.072-3.273-6.068-6.757-8.956-10.45c-0.252-0.327-0.63-0.504-0.923-0.789C147.674%2C297.682%2C134.295%2C264.419%2C134.295%2C226.623z%20%20%20%20%20%20M160.18%2C359.147l3.769-15.058c0.017%2C0.025%2C0.042%2C0.042%2C0.059%2C0.059c20.287%2C22.066%2C46.097%2C37.611%2C75.188%2C41.145%20%20%20%20%20c0.017%2C0%2C0.034%2C0.008%2C0.05%2C0.008l-22.78%2C51.242L160.18%2C359.147z%20M220.084%2C486.811l5.028-25.172h53.391l2.132%2C10.635l2.896%2C14.479%20%20%20%20%20L220.084%2C486.811z%20M287.148%2C436.543l-22.78-51.242c0.017%2C0%2C0.034-0.008%2C0.05-0.008c29.092-3.534%2C54.902-19.078%2C75.188-41.145%20%20%20%20%20c0.017-0.017%2C0.042-0.034%2C0.059-0.059l3.769%2C15.058L287.148%2C436.543z%22%20data-original%3D%22%23000000%22%20class%3D%22%22%20fill%3D%22%23A9a9a9%22/%3E%0A%09%09%09%3Cpath%20d%3D%22M331.617%2C43.415l-10.5%2C55.111l57.436-11.046C366.626%2C69.762%2C350.679%2C54.746%2C331.617%2C43.415z%22%20data-original%3D%22%23000000%22%20class%3D%22%22%20fill%3D%22%23a9a9a9%22/%3E%0A%09%09%09%3Cpath%20d%3D%22M171.996%2C43.415c-19.062%2C11.331-35.009%2C26.347-46.936%2C44.066l57.436%2C11.046L171.996%2C43.415z%22%20data-original%3D%22%23000000%22%20class%3D%22%22%20fill%3D%22%23a9a9a9%22/%3E%0A%09%09%3C/g%3E%0A%09%3C/g%3E%0A%3C/g%3E%3C/g%3E%20%3C/svg%3E")!important;
}

.fa-twitter:before {
    content: url('data:image/svg+xml,%3Csvg%20xmlns%3D%22http%3A//www.w3.org/2000/svg%22%20xmlns%3Axlink%3D%22http%3A//www.w3.org/1999/xlink%22%20version%3D%221.1%22%20id%3D%22Capa_1%22%20x%3D%220px%22%20y%3D%220px%22%20viewBox%3D%220%200%20512%20512%22%20style%3D%22enable-background%3Anew%200%200%20512%20512%3B%22%20xml%3Aspace%3D%22preserve%22%20width%3D%2230px%22%20height%3D%2230px%22%20class%3D%22%22%3E%3Cg%3E%3Cg%3E%0A%09%3Cg%3E%0A%09%09%3Cpath%20d%3D%22M511.68%2C498.731l-16.576-66.283c-3.797-15.125-13.952-27.819-27.904-34.773l-110.443-55.211%20%20%20%20c-2.603-1.301-5.675-1.472-8.448-0.469c-2.773%2C1.003-4.971%2C3.115-6.144%2C5.824l-54.933%2C128.192l-8.085-16.149%20%20%20%20c7.147-9.792%2C19.52-28.971%2C19.52-43.861c0-19.733-16.363-32-42.667-32s-42.667%2C12.267-42.667%2C32%20%20%20%20c0%2C14.891%2C12.373%2C34.069%2C19.499%2C43.84l-8.085%2C16.149l-54.933-128.192c-1.173-2.709-3.371-4.821-6.144-5.824%20%20%20%20c-2.773-0.981-5.803-0.832-8.448%2C0.469L44.8%2C397.675c-13.931%2C6.976-24.107%2C19.648-27.904%2C34.773L0.32%2C498.731%20%20%20%20c-0.789%2C3.179-0.085%2C6.571%2C1.941%2C9.152c2.005%2C2.603%2C5.12%2C4.117%2C8.405%2C4.117h490.667c3.285%2C0%2C6.4-1.515%2C8.427-4.096%20%20%20%20C511.765%2C505.323%2C512.469%2C501.931%2C511.68%2C498.731z%22%20data-original%3D%22%23000000%22%20class%3D%22%22%20fill%3D%22%23a9a9a9%22/%3E%0A%09%3C/g%3E%0A%3C/g%3E%3Cg%3E%0A%09%3Cg%3E%0A%09%09%3Cpath%20d%3D%22M401.173%2C175.637c-1.856-5.589-7.893-8.619-13.483-6.741c-5.589%2C1.856-8.597%2C7.893-6.741%2C13.483%20%20%20%20c2.005%2C5.973%2C3.051%2C12.992%2C3.051%2C20.288c0%2C20.629-7.68%2C31.979-11.349%2C31.509c-2.88-1.792-6.485-2.112-9.643-0.832%20%20%20%20s-5.504%2C4.011-6.336%2C7.317C341.738%2C300.885%2C301.29%2C341.333%2C256%2C341.333s-85.739-40.448-100.672-100.672%20%20%20%20c-0.832-3.307-3.52-5.781-6.677-7.061c-3.179-1.301-7.083-0.725-9.984%2C1.067c-2.667%2C0-10.667-11.349-10.667-32%20%20%20%20c0-7.296%2C1.045-14.315%2C3.051-20.288c1.856-5.568-1.152-11.627-6.763-13.483c-5.589-1.877-11.627%2C1.152-13.483%2C6.741%20%20%20%20c-2.709%2C8.107-4.139%2C17.451-4.139%2C27.029c0%2C29.717%2C13.141%2C52.288%2C30.741%2C53.312C156.778%2C320.213%2C203.456%2C362.667%2C256%2C362.667%20%20%20%20s99.221-42.453%2C118.592-106.688c17.6-1.024%2C30.741-23.595%2C30.741-53.312C405.333%2C193.088%2C403.904%2C183.744%2C401.173%2C175.637z%22%20data-original%3D%22%23000000%22%20class%3D%22%22%20fill%3D%22%23a9a9a9%22/%3E%0A%09%3C/g%3E%0A%3C/g%3E%3Cg%3E%0A%09%3Cg%3E%0A%09%09%3Cpath%20d%3D%22M256%2C0c-82.347%2C0-149.333%2C66.987-149.333%2C149.333c0%2C9.899%2C1.216%2C20.053%2C3.84%2C31.979c1.067%2C4.885%2C5.397%2C8.384%2C10.411%2C8.384%20%20%20%20c12.075%2C0%2C26.667%2C2.176%2C26.816%2C2.197c0.619%2C0.085%2C1.152%2C0.128%2C1.813%2C0.107c5.824%2C0.085%2C10.88-4.693%2C10.88-10.667%20%20%20%20c0-1.301-0.235-2.517-0.64-3.669c-1.301-21.291%2C0.043-42.816%2C3.008-49.664h1.621c36.885%2C0%2C55.275-11.712%2C63.765-26.496%20%20%20%20c22.997%2C22.699%2C69.248%2C44.587%2C113.963%2C47.488c2.539%2C17.92%2C10.389%2C35.755%2C10.795%2C36.651c1.707%2C3.904%2C5.547%2C6.357%2C9.728%2C6.357%20%20%20%20c0.277%2C0%2C0.576%2C0%2C0.875-0.021l28.416-2.325c4.672-0.384%2C8.533-3.776%2C9.557-8.341c2.603-11.925%2C3.819-22.08%2C3.819-31.979%20%20%20%20C405.333%2C66.987%2C338.346%2C0%2C256%2C0z%22%20data-original%3D%22%23000000%22%20class%3D%22%22%20fill%3D%22%23a9a9a9%22/%3E%0A%09%3C/g%3E%0A%3C/g%3E%3Cg%3E%0A%09%3Cg%3E%0A%09%09%3Cpath%20d%3D%22M224%2C192h-21.333c-5.888%2C0-10.667%2C4.779-10.667%2C10.667s4.779%2C10.667%2C10.667%2C10.667H224%20%20%20%20c5.888%2C0%2C10.667-4.779%2C10.667-10.667S229.888%2C192%2C224%2C192z%22%20data-original%3D%22%23000000%22%20class%3D%22%22%20fill%3D%22%23a9a9a9%22/%3E%0A%09%3C/g%3E%0A%3C/g%3E%3Cg%3E%0A%09%3Cg%3E%0A%09%09%3Cpath%20d%3D%22M309.333%2C192H288c-5.888%2C0-10.667%2C4.779-10.667%2C10.667s4.779%2C10.667%2C10.667%2C10.667h21.333%20%20%20%20c5.888%2C0%2C10.667-4.779%2C10.667-10.667S315.221%2C192%2C309.333%2C192z%22%20data-original%3D%22%23000000%22%20class%3D%22%22%20fill%3D%22%23a9a9a9%22/%3E%0A%09%3C/g%3E%0A%3C/g%3E%3C/g%3E%20%3C/svg%3E%20')!important;
}

.q-chip {
	margin: 1px;
}

.text-primary {
	color:#42b983!important;
}

.text-warn {
	color: #F2BC37;
}

.text-danger {
	color: #F23C37;
}
</style>