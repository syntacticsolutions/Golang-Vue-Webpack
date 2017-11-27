<template>
    <div class="panel">
        <div class="project-item" v-for="(item, id) in projects" :key="id">
            <span class="project-title">
                {{ item.title }}
            </span>
            <span class="project-dates">
                {{ showDates(item.start_date.String, item.end_date.String) }}
            </span>
            <ul>
                <li v-for=""></li>
            </ul>
        </div>
    </div>
</template>

<script>
const _ = {
    each: require('lodash/each'),
    isEmpty: require('lodash/isEmpty')
}

export default {
    data() {
        return {
            projects: null,
            markers: null,
            marker_types: null
        }
    },
    methods: {
        regenerateProjects(){
            if(!this.projects ||
            !this.markers ||
            !this.marker_types)return;
            
            this.project['markers'] = {}
            // && !_.isEmpty(this.contractors))
            _.each(this.markers, (marker)=>{
                this.project.markers[marker.id] = marker
            })

        },
        showDates(start, end){
            return this.userFormatDate(start) + ' - ' + this.userFormatDate(end)
        },
        userFormatDate(date){
            let arr =  date.split(' ')[0].split('-')
            return arr[1] + '/' + arr[2] + '/' + arr[0]
        }
    },  
    mounted(){
        axios.get(config.host + '/api/projects')
        .then(res => {
            this.projects = {}
            _.each(res.data.projects, (project)=>{
                this.projects[project.id] = project;
            })
            this.regenerateProjects()
        })

        axios.get(config.host + '/api/markers')
        .then(res => {
            this.markers = {}
            _.each(res.data.markers, (marker)=>{
                this.markers[marker.id] = marker;
            })
            this.regenerateProjects()
        })

        axios.get(config.host + '/api/marker_types')
        .then(res => {
            this.marker_types = {}
            _.each(res.data.marker_types, (marker_type)=>{
                this.marker_types[marker_type.id] = marker_type;
            })
            this.regenerateProjects()
        })

        window.panel = this;
    }
}
</script>

<style lang="scss" scoped>

.panel {
    height: 93vh;
    width: 350px;
    background-color: white;
    z-index: 1;
    position: absolute;
}

.project-item {
    background-color:red;
    width:100%;
    height: 125px;
    float:left;
    clear:left;
}

.project-title{
    padding-left: 10px;
    margin-top: 15px;
    font-size: 24px;
    clear: left;
    float:left;
    top: 18px;
}

.project-dates {
    float: left;
    clear: left;
    padding: 3px 10px;
}

ul {
    list-style-type: none;
}

</style>
