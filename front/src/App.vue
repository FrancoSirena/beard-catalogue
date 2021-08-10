<template>
  <div id="app">
    <button type="button" @click="getSome()">API Call</button>
    <div v-if="Something != null">
      <h1>{{ Something }}</h1>
    </div>
  </div>
</template>

<script>

export default {
  name: "App",
  components: {},
  data() {
    return {
      Something: null,
    };
  },
  methods: {
    getSome() {
      this.axios.get("/try").then(({ data }) => {
        console.log(data)
        this.Something = data
      });
    },
  },
};
</script>

