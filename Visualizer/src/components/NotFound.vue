<script setup lang="ts">
import { defineModel } from "vue";

import store from "../Store";

import notFoundIcon from "/Icons/notFound.png";

const isEmailError = defineModel<boolean>("isEmailError");
</script>

<template>
  <div class="flex flex-col mx-[auto] block w-[fit-content]">
    <img
      class="w-[100px] mx-[auto]"
      :src="notFoundIcon"
      alt="404 error when searching"
    />
    <p class="">
      {{ isEmailError ? "The Email" : `The word: ${store.searchedValue}` }} was
      not found
    </p>
  </div>
</template>
