<template>
  <div class="container row">
    <div class="row">
      <div class="row">
        <form class="row" name="form" @submit.prevent>
          <div>
            <label for="form-name">名前</label>
            <input id="form-name" v-model="name" type="text" name="name">
          </div>

          <div>
            <label for="form-password">パスワード</label>
            <input id="form-password" v-model="password" type="password" name="password">
          </div>
          <div>
            <button id="form__save" class="btn btn-dark" @click="submit">
              ログイン
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, useRouter, ref, useContext } from '@nuxtjs/composition-api'
export default defineComponent({
  name: 'TopPage',
  setup () {
    const { $axios } = useContext()
    const router = useRouter()
    const name = ref('')
    const password = ref('')
    const submit = async (event: any) => {
      event.preventDefault()
      const params = { name: name.value, password: password.value }
      try {
        await $axios.post('/api/login', params)
        router.push('/')
      } catch (err) {
        // console.log(err)
      }
    }

    return {
      name,
      password,
      submit
    }
  },
  head: {}
})
</script>
<style lang="scss" scoped>
</style>
