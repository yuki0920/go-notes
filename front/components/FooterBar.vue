<template>
  <footer>
    <div class="p-1">
      <ul class="list-unstyled d-flex justify-content-center flex-wrap mb-2">
        <li class="p-1">
          <a href="https://twitter.com/YukiWebTech" target="_blank" rel="noopener">Twitter</a>
        </li>
        <li class="p-1">
          <a href="https://github.com/yuki0920" target="_blank" rel="noopener">GitHub</a>
        </li>
        <template v-if="isAuthenticated">
          <li class="p-1">
            <nuxt-link to="/articles/new">
              New Article
            </nuxt-link>
          </li>
          <li class="p-1">
            <a href="javascript:void(0)" @click="logout">Log out</a>
          </li>
        </template>
      </ul>
      <div class="text-center text-black-50">
        © 2021 Go Notes
      </div>
    </div>
  </footer>
</template>

<script lang="ts">
import { defineComponent, useContext, useRouter, ref } from '@nuxtjs/composition-api'
export default defineComponent({
  name: 'FooterBar',
  setup () {
    const { $axios } = useContext()
    const router = useRouter()

    const isAuthenticated = ref<Boolean>(false)
    const auth = async () => {
      const { data } = await $axios.get('/api/auth')
      isAuthenticated.value = data.IsAuthenticated
    }
    auth()

    const logout = async () => {
      try {
        await $axios.post('/api/logout')
        router.push('/')
      } catch (err) {
        // console.error(err)
      }
    }

    return {
      isAuthenticated,
      logout
    }
  }
})
</script>

<style lang="scss" scoped>
</style>
