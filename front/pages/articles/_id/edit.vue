<template>
  <div v-if="article" class="l-col l-row l-v-padd">
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
              <nuxt-link :to="`../${article.id}`">
                キャンセル
              </nuxt-link>
            </a>
            <a id="article-form__save" href="javascript:void(0)" class="btn btn-dark" @click="submit">
              保存
            </a>
            <a href="javascript:void(0)" class="btn" @click="deleteArticle">
              <b-icon-trash style="font-size: 2rem; color: red;" />
            </a>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, useRoute, useRouter, useContext, ref } from '@nuxtjs/composition-api'
import { Article } from '~/types/article'
export default defineComponent({
  name: 'Articles',
  setup () {
    const route = useRoute()
    const id = route.value.params.id
    const { $axios } = useContext()
    const router = useRouter()

    const isAuthenticated = ref<Boolean>(false)
    const article = ref<Article | null>(null)

    onMounted(async () => {
      const { data: articleData } = await $axios.get(`/api/articles/${id}`)
      if (articleData) { article.value = articleData }

      const { data: authData } = await $axios.get('/api/auth')
      isAuthenticated.value = authData.IsAuthenticated
      if (isAuthenticated.value === false) { router.push(`../${article.value?.id}`) }
    })

    const deleteArticle = async () => {
      try {
        await $axios.delete(`/api/articles/${id}`)
        router.push('../../')
      } catch (err) {
        // console.error(err)
      }
    }

    const submit = async (event: any) => {
      event.preventDefault()
      const params = { title: article.value?.title, body: article.value?.body }
      try {
        await $axios.put(`/api/articles/${id}`, params)
        router.push(`../${article.value?.id}`)
      } catch (err) {
        // console.log(err)
      }
    }

    return {
      article,
      submit,
      deleteArticle
    }
  }
})
</script>

<style lang='scss' scoped>
</style>
