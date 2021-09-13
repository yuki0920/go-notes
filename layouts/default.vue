<template>
    <header>
      <div class="header">
        <a class="header__logo" href="/">My Tech Blog</a>
      </div>
    </header>
    <nuxt />
    <footer>
      <div class="footer">
        <ul class="footer__menu">
          <li class="footer__menu-item"><a href="">Twitter</a></li>
          <li class="footer__menu-item"><a href="">GitHub</a></li>
          <li class="footer__menu-item"><a href="">プライバシーポリシー</a></li>
          <li class="footer__menu-item"><a href="">運営者</a></li>
        </ul>
        <div class="footer__copy-right">
          © 2021 Go Blog
        </div>
      </div>
    </footer>
</template>

<script lang="ts">
import { defineComponent } from '@nuxtjs/composition-api'

export default defineComponent({
  setup () {
  }
})
</script>
<style lang="scss">
</style>
