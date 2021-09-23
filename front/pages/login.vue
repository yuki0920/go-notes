<template>
  <div class="l-col l-row l-v-padd">
    <div class="article-new l-row">
      <div class="article-new__form l-row">
        <!-- TODO: 入力フォームと共通化したい -->
        <form class="article-form l-row" name="article-form" @submit.prevent>
          <!-- TODO: TODO: エラーを表示したい -->
          <div class="article-form__title">
            <label class="article-form__label" for="form-name">名前</label>
            <input id="form-name" v-model="name" class="article-form__input" type="text" name="name">
          </div>

          <div class="article-form__body">
            <label class="article-form__label" for="form-password">パスワード</label>
            <input id="form-password" v-model="password" class="article-form__input" type="password" name="password">
          </div>
          <div class="article-form__footer">
            <button id="article-form__save" class="article-form__save btn--primary" @click="submit">
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
