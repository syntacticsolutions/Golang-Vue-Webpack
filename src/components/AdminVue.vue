<template>
    <div class="appContainer">
        <q-toolbar color="orange">
            <!-- <q-btn flat><q-icon size="50px" name="mail" /></q-btn> -->
            <q-toolbar-title>
            Admin Panel
            </q-toolbar-title>
            <!-- <q-btn flat><q-icon size="50px" name="alarm" /></q-btn>
            <q-btn flat><q-icon size="50px" name="router" /></q-btn> -->
            <q-btn flat><q-icon size="50px" name="menu" /></q-btn>
        </q-toolbar>
        <Panel @makeMarker="makeMarker"></Panel>
        <Map :markers="markers"></Map>
        <!-- <Toolbar></Toolbar> -->
    </div>
</template>

<script>

import { QToolbar, QToolbarTitle, QIcon, QBtn } from 'quasar'
import Map from './Map.vue'
import Toolbar from './Toolbar.vue'
import Panel from './Panel.vue'

export default {
    data() {
        return {
            markers:{}
        }
    },
    components: {
        QToolbar,
        QToolbarTitle,
        QIcon,
        QBtn,
        Toolbar,
        Panel,
        Map
    },
    methods: {
        makeMarker(marker) {
            var img = {
              url: 'data:image/svg+xml;base64,' + btoa(marker.svg),
              size: new google.maps.Size(192, 21),
              scaledSize: new google.maps.Size(192,21),
              origin: new google.maps.Point(0, 0),
              anchor: new google.maps.Point(88, 53)
          };

            this.markers[marker.id] = new google.maps.Marker({
                position: {lat: marker.lat, lng: marker.lng},
                map:map,
                icon: img,
                zIndex: 2,
                optimized: false
            })
        }
    },
    mounted() {
        window.view = this;
    }
  
}
</script>

<style lang="scss" scoped>

</style>
