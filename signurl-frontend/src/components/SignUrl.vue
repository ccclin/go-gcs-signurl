<template>
  <div>
    <div class="row">
      <div class="row col offset-l3 l6 offset-m1 m10 s12">
        <a href="https://www.mile.cloud/" target="_blank"><img src="https://www.mile.cloud/wp-content/uploads/2018/08/CloudMile-300x119.png" alt="CloudMile"/></a>
        <h4>GCS SignURL with MP4 Video</h4>
      </div>
      <div class="row col offset-l3 l6 offset-m1 m10 s12">
        <video width="665" controls :src="signURL"></video>
      </div>
      <div class="row col offset-l3 l6 offset-m1 m10 s12">
        <div class="card">
          <div class="card-content">
            <div v-if="uploading" class="progress">
              <div class="indeterminate"></div>
            </div>
            <div v-else >
              <label>File
                <input type="file" id="file" ref="file" v-on:change="handleFileUpload()"/>
              </label>
              <div class="right-align">
                <button class="btn waves-effect waves-light btn" v-bind:class="submitClass" v-on:click="submitFile()">Submit</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SignUrl',

  data () {
    return {
      signURL: '',
      postSignURL: '',
      file: {},
      uploading: false
    }
  },

  computed: {
    submitClass: function () {
      return Object.entries(this.file).length === 0 && this.file.constructor === Object ? 'disabled' : ''
    }
  },

  methods: {
    getSignURL: function () {
      var self = this
      fetch(process.env.VUE_APP_API_HOST + 'api/v1/sign_url', {
        method: "GET"
      }).then(function(response) {
        return response.json()
      }).then(function(object) {
        self.signURL = object.signURL
        self.postSignURL = object.postSignURL
      })
    },

    handleFileUpload(){
      this.file = this.$refs.file.files[0];
    },

    submitFile(){
      var file = this.file
      var self = this
      this.uploading = true
      this.file = {}
      fetch(this.postSignURL, {
        method: 'PUT',
        body: file
      }).then(function() {
        self.getSignURL()
        self.uploading = false
      })
    }
  },

  created () {
    this.getSignURL()
  }
}
</script>
