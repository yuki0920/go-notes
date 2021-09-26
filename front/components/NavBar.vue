<template>
  <div class="bg-dark p-2">
    <nuxt-link class="navbar-brand text-white" to="/">
      Go Notes
    </nuxt-link>
  </div>
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api'
export default defineComponent({
  name: 'NavBar'
})
</script>

<style lang="scss" scoped>
</style>
