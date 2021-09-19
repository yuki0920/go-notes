<template>
  <div class="header">
    <nuxt-link class="header__logo" to="/">
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
.header {
  background: #222;
  padding: 16px;
  color: #fff;
}

.header__logo {
  color: #fff;
}
</style>
