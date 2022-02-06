<template>
  <div v-if="article" class="container">
    <form>
      <!-- TODO: TODO: エラーを表示したい -->
      <div class="form-group">
        <label for="form-title">タイトル</label>
        <input
          id="form-title"
          v-model="article.title"
          class="form-control"
          type="text"
          name="title"
        >
      </div>
      <div>
        <multiselect
          v-model="selectedTags"
          tag-placeholder="Add this as new tag"
          placeholder="Search or add a tag"
          label="title"
          track-by="id"
          :options="categoryOptions"
          :multiple="true"
          :taggable="true"
          @tag="addTag"
        />
      </div>
      <div class="form-group">
        <label for="form-body">本文</label>
        <textarea
          id="form-body"
          v-model="article.body"
          class="form-control"
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
</template>

<script lang="ts">
import { defineComponent, onMounted, useRoute, useRouter, useContext, ref } from '@nuxtjs/composition-api'
import { Article } from '~/types/article'
import { Category } from '~/types/category'
export default defineComponent({
  name: 'ArticleEdit',
  setup () {
    const route = useRoute()
    const id = route.value.params.id
    const { $axios } = useContext()
    const router = useRouter()

    const isAuthenticated = ref<Boolean>(false)
    const article = ref<Article | null>(null)
    type categoryOption = {
      id: number,
      title: string,
    }
    const categoryOptions = ref<categoryOption[]>([])
    const selectedTags = ref<categoryOption[]>([])

    onMounted(async () => {
      const { data: articleData } = await $axios.get(`/api/articles/${id}`)
      if (articleData) { article.value = articleData }
      article.value?.categories?.forEach((category: Category) => {
        selectedTags.value.push({
          id: category.id,
          title: category.title
        })
      })

      const { data: categoryData } = await $axios.get('/api/categories')
      if (categoryData) {
        categoryData.forEach((category: Category) => {
          categoryOptions.value.push({ id: category.id, title: category.title })
        })
      }

      const { data: authData } = await $axios.get('/api/auth')
      isAuthenticated.value = authData.IsAuthenticated
      if (isAuthenticated.value === false) { router.push(`../${article.value?.id}`) }
    })

    const addTag = async (newTag: string) => {
      const { data: id } = await $axios.post('/api/categories', { title: newTag })
      const tag = {
        id: parseInt(id, 10),
        title: newTag
      }
      categoryOptions.value.push(tag)
      selectedTags.value.push(tag)
    }

    const deleteArticle = async () => {
      await $axios.delete(`/api/articles/${id}`)
      router.push('../../')
    }

    const submit = async (event: any) => {
      event.preventDefault()
      const params = {
        title: article.value?.title,
        body: article.value?.body,
        category_ids: selectedTags.value.map(tag => tag.id)
      }
      await $axios.put(`/api/articles/${id}`, params)
      router.push(`../${article.value?.id}`)
    }

    return {
      article,
      submit,
      deleteArticle,
      categoryOptions,
      selectedTags,
      addTag
    }
  }
})
</script>

<style lang='scss' scoped>
</style>
