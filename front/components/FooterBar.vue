<template>
  <footer>
    <div class="footer">
      <ul class="footer__menu">
        <li class="footer__menu-item">
          <a href="https://twitter.com/YukiWebTech" target="_blank" rel="noopener">Twitter</a>
        </li>
        <li class="footer__menu-item">
          <a href="https://github.com/yuki0920" target="_blank" rel="noopener">GitHub</a>
        </li>
        <li class="footer__menu-item">
          <nuxt-link to="/login">Log in</nuxt-link>
        </li>
        <li class="footer__menu-item">
          <a href="javascript:void(0)" @click="logout">Log out</a>
        </li>
      </ul>
      <div class="footer__copy-right">
        © 2021 Go Notes
      </div>
    </div>
  </footer>
</template>

<script lang="ts">
import { defineComponent, useContext } from '@nuxtjs/composition-api'
export default defineComponent({
  name: 'FooterBar',
  setup () {
    const { $axios } = useContext()
    const logout = async () => {
      try {
        await $axios.post('/api/logout')
      } catch (err) {
        // console.error(err)
      }
    }

    return {
      logout
    }
  }
})
</script>

<style lang="scss" scoped>
.footer {
  background: #fff;
  padding: 4px 10%;
  color: #222;
  font-size: 16px;
}

.footer__menu {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
}

.footer__menu-item {
  list-style: none;
  padding: 2px 4px;
}

.footer__copy-right {
  text-align: center;
  color: #bbb;
}
</style>
