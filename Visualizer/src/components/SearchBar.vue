<script setup lang="ts">
import { defineModel } from "vue";

import chevronIcon from "/Icons/chevron.png";
import searchIcon from "/Icons/search.png";
import store from "../Store";

const searchValue = defineModel<string | undefined | null>("searchTerm");
const SearchFunction = defineModel<() => Promise<void>>("SearchFunction");
const isLoading = defineModel<boolean>("isLoading");

const handleChangeOrder = () => {
  store.setOrder(!store.order);
  SearchFunction.value && SearchFunction.value();
};
</script>

<template>
  <div
    class="flex flex-row-reverse justify-center items-center gap-4 mt-4"
  >
    <!-- Button to toggle order -->
    <button
      :onclick="handleChangeOrder"
      class="bg-[#686868] hover:bg-[#4f4f4f] text-white font-medium rounded-full p-3 transition-transform duration-300 transform hover:scale-105 active:scale-110 flex items-center gap-2"
    >
      <span class="text-sm">Date</span>
      <img
        :class="`w-[18px] transition-transform duration-300 transform ${
          store.order ? 'rotate-180' : 'rotate-0'
        }`"
        :src="chevronIcon"
        alt="chevron for order"
      />
    </button>

    <!-- Search input -->
    <div class="relative w-[60%] min-w-[250px] max-w-[600px]">
      <input
        class="w-full bg-[#ffffff] text-black rounded-md text-lg p-3 pl-5 pr-12 focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all"
        type="text"
        placeholder="Search..."
        v-model="searchValue"
        @keydown.enter="SearchFunction"
        name="search"
        id="search"
        minlength="2"
        :disabled="isLoading"
      />
      <label for="search" class="absolute top-1/2 right-4 transform -translate-y-1/2 cursor-pointer">
        <button v-on:click="SearchFunction">
          <img class="w-6 h-6" :src="searchIcon" alt="search button" />
        </button>
      </label>
    </div>
  </div>
</template>