<template>
  <div class="l-col l-row l-v-padd">
    <div class="l-row">
      <div class="l-row">
        <!-- TODO: 入力フォームと共通化したい -->
        <form class="l-row">
          <!-- TODO: TODO: エラーを表示したい -->
          <div>
            <label class="d-block" for="form-title">タイトル</label>
            <input id="form-title" v-model=" article.title" class="w-100" type="text" name="title">
          </div>

          <div class="">
            <label class="d-block" for="form-body">本文</label>
            <textarea
              id="form-body"
              v-model="article.body"
              class="w-100"
              rows="20"
              name="body"
            />
          </div>

          <div>
            <a href="javascript:void(0)" class="btn btn-secondary">
              <nuxt-link to="/">
                キャンセル
              </nuxt-link>
            </a>
            <a href="javascript:void(0)" class="btn btn-dark" @click="submit">
              保存
            </a>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, useRouter, useContext, reactive } from '@nuxtjs/composition-api'

export default defineComponent({
  name: 'ArticleNew',
  setup () {
    const { $axios } = useContext()
    const article = reactive({ title: null, body: null })
    const router = useRouter()

    const submit = async (event: any) => {
      event.preventDefault()
      const params = { title: article.title, body: article.body }
      try {
        const { data } = await $axios.post('/api/articles', params)
        const id = data.Article.id
        router.push(`/articles/${id}`)
      } catch (err) {
        // console.log(err)
      }
    }

    return {
      article,
      submit
    }
  }
})
</script>

<style lang='scss' scoped>
</style>
