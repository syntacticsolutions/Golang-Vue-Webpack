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
                <li 
                v-if="marker_types"
                v-for="(marker, id) in markers" 
                :key="id" 
                v-html="marker.svg"></li>
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
            marker_types: null,
        }
    },
    methods: {
        regenerateProjects(){
            if(!_.isEmpty(this.projects)
            && !_.isEmpty(this.markers)
            && !_.isEmpty(this.marker_types)
            && !_.isEmpty(this.contractors))

            for(var id in this.markers){
                this.markers[id]['svg'] = this.marker_types[this.markers[id].type_id].svg.replaceAll(
                    '{{}}', 
                    this.contractors[this.projects[this.markers[id].project_id].contractor_id].color
                )
                this.$emit('makeMarker', this.markers[id]);
                this.projects[this.markers[id].project_id].markers[id] = this.markers[id]
            }

        },
        showDates(start, end){
            let first = start ? this.userFormatDate(start) : 'unknown'
            let last = end ? this.userFormatDate(end) : 'unknown'
            
            return first + ' - ' + last;
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
                project['markers'] = {}
                this.projects[project.id] = project;
            })
            this.regenerateProjects()
        })

        axios.get(config.host + '/api/markers')
        .then(res => {
            this.markers = {}
            _.each(res.data.markers, (marker)=>{
                marker.svg = ""
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

        axios.get(config.host + '/api/contractors')
        .then(res => {
            this.contractors = {}
            _.each(res.data.contractors, (contractor)=>{
                this.contractors[contractor.id] = contractor;
            })
        })

        window.panel = this;
    }
}
String.prototype.replaceAll = function(delimiter, replacement){
    return this.split(delimiter).join(replacement);
}

</script>

<style lang="scss" scoped>

.panel {
    height: 93vh;
    width: 350px;
    background-color: white;
    z-index: 1;
    position: absolute;
    overflow-y: scroll;
}

.project-item {
    background-color:white;
    width:100%;
    height: 125px;
    float:left;
    clear:left;
}

.project-item:hover{
    color:white;
    background-color:lightgrey;
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
    display: inline-flex;
    float:left;
    clear:left;
    padding-left: 10px;
    margin-top: 5px;
}

</style>
