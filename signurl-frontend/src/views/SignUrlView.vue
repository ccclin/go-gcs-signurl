<template>
  <div>
    <div class="row">
      <div class="row col offset-l3 l6 offset-m1 m10 s12">
        <CustomLogo />
      </div>
      <div class="row col offset-l3 l6 offset-m1 m10 s12">
        <h4>GCS SignURL with MP4 Video</h4>
      </div>
      <div class="row col offset-l3 l6 offset-m1 m10 s12">
        <SignUrl />
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import SignUrl from '@/components/SignUrl.vue'
import CustomLogo from '@/components/CustomLogo.vue'

export default defineComponent({
  name: 'SignUrlView',
  components: {
    SignUrl,
    CustomLogo
  }
})
</script>

<style></style>
