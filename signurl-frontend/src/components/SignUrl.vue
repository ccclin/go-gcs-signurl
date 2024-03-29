<template>
  <div>
    <div class="row col">
      <ul class="collection">
        <li class="collection-item">
          <span class="title blue-text text-darken-2">SignURL For GET</span>
          <div style="width:840px;overflow:auto"><span class="red-text text-darken-4">{{ state.signURL }}</span></div>
        </li>
        <li class="collection-item">
          <span class="title blue-text text-darken-2">SignURL For PUT</span>
          <div style="width:840px;overflow:auto"><span class="red-text text-darken-4">{{ state.postSignURL }}</span></div>
        </li>
      </ul>
    </div>
    <div class="row col">
      <video width="840" controls :src="state.signURL"></video>
    </div>
    <div class="row col l12">
      <div class="card">
        <div class="card-content">
          <div v-if="state.uploading" class="progress">
            <div class="indeterminate"></div>
          </div>
          <div v-else>
            <label
              >File
              <input type="file" id="file" ref="files" v-on:change="handleFileUpload()" />
            </label>
            <div class="right-align">
              <button
                class="btn waves-effect waves-light btn"
                v-bind:class="submitClass()"
                v-on:click="submitFile()"
              >
                Submit
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { defineComponent, reactive, onMounted, ref } from 'vue'
export default defineComponent({
  name: 'SignUrl',
  setup() {
    const files = ref(null)
    const state = reactive({
      signURL: '',
      postSignURL: '',
      uploading: false,
      file: {}
    })

    const getSignURL = async () => {
      fetch(import.meta.env.VITE_APP_API_HOST + 'api/v1/sign_url', {
        method: 'GET'
      })
        .then(function (response) {
          return response.json()
        })
        .then(function (object) {
          state.signURL = object.signURL
          state.postSignURL = object.postSignURL
        })
    }

    const handleFileUpload = () => {
      state.file = files.value.files[0]
    }

    const submitFile = () => {
      let file = state.file
      state.uploading = true
      fetch(state.postSignURL, {
        method: 'PUT',
        body: file
      }).then(async function () {
        await getSignURL()
        state.uploading = false
        state.file = {}
      })
    }

    const submitClass = () => {
      return state.file['type'] === 'video/mp4' ? '' : 'disabled'
    }

    onMounted(async () => {
      await getSignURL()
    })

    return {
      state,
      files,
      handleFileUpload,
      submitFile,
      submitClass
    }
  }
})
</script>

<style></style>
