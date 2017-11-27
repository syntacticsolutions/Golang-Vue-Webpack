<template>
  <div id="map"></div>
</template>

<script>
export default {
    data() {
        return {
            api_key: "AIzaSyCOSu1gCWOJNX8tuZHnBL5a-NkXc_orSfY"
        }
    },
    mounted() {
        let newScript = document.createElement("script");
        newScript.src = "https://maps.googleapis.com/maps/api/js?key=" + this.api_key + "&callback=initMap";
        $('body')[0].appendChild(newScript);
    }
}

window.initMap = () => {
    let location = {lat: -25.363, lng: 131.044};
    window.map = new google.maps.Map(document.getElementById('map'), {
        zoom: 4,
        center: location
    });
}
</script>

<style lang="scss" scoped>
      #map {
        height: 93vh;
        width: 100%;
       }
</style>
