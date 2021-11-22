<template>
  <div>
    <nav-bar />
    <nuxt class="main-content" />
    <footer-bar />
  </div>
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api'

export default defineComponent({
  name: 'DefaultLayout',
  setup () {
  }
})
</script>

<style lang="scss">
.main-content {
  // NOTE: ヘッダー56px、フッター70pxを引いている
  min-height: calc(100vh - 56px - 70px);
}
</style>
