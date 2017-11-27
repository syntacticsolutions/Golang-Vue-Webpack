<template>
  <div id="map"></div>
</template>

<script>
export default {
    props:['markers'],
    data() {
        return {
            api_key: "AIzaSyCOSu1gCWOJNX8tuZHnBL5a-NkXc_orSfY"
        }
    },
    methods: {
        initMap() {
            let location = 
            {
                "lng": -121.28925,
                "lat": 38.761578
            }
            window.map = new google.maps.Map(document.getElementById('map'), {
                zoom: 15,
                center: location
            });
        }
    },
    mounted() {

        this.initMap();

    }
}
</script>

<style lang="scss" scoped>
      #map {
        height: 93vh;
        width: 100%;
       }
</style>
