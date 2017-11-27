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
        <Panel></Panel>
        <Map></Map>
        <Toolbar></Toolbar>
    </div>
</template>

<script>

import { QToolbar, QToolbarTitle, QIcon, QBtn } from 'quasar'
import Map from './Map.vue'
import Toolbar from './Toolbar.vue'
import Panel from './Panel.vue'

export default {
    components: {
        QToolbar,
        QToolbarTitle,
        QIcon,
        QBtn,
        Toolbar,
        Panel,
        Map
    }
  
}
</script>

<style lang="scss" scoped>

</style>
