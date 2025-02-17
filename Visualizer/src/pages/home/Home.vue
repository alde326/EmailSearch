<script setup lang="ts">
import { ref, watch, computed, onMounted } from "vue";
import { useRouter } from "vue-router";
import store from "../../Store";
import SearchBar from "../../components/SearchBar.vue";
import Loader from "../../components/Loader.vue";
import NotFound from "../../components/NotFound.vue";
import search from "../../helpers/fetchEmails";
import { Hits, Hit } from "../../types/emailTypes";
import { PAGINATION_VALUE } from "../../constants/constants";

const searchValue = ref<string>(store.searchedValue);
const isLoading = ref<boolean>(false);
const notFound = ref<boolean>(false);
const selectedEmail = ref<Hit | null>(null); // Almacena el email seleccionado

const router = useRouter();

const hasResults = computed(
  () => !isLoading.value && Array.isArray(store.fetchedEmails) && store.fetchedEmails.length > 0
);

const handleResetPagination = () => {
  store.setHasMoreResults(true);
  store.resetPagination();
};

const handleSearch = async () => {
  let isUserSearching = true;
  if (!searchValue.value || searchValue.value.length < 2) {
    store.setIsEmailResponseForPagination(true);
    isLoading.value = false;
    notFound.value = false;
    isUserSearching = false;
  }

  store.setIsSearching(true);
  isLoading.value = true;
  const oldSearchedTerm = store.searchedValue;
  store.setSearchedValue(searchValue.value);

  if (searchValue.value !== oldSearchedTerm) handleResetPagination();

  const hitsInformation: Hits | undefined = await search(
    searchValue.value,
    String(store.pagination),
    store.order ? "-" : "",
    isUserSearching
  );

  let allResults: Hit[] | undefined | null = [];
  if (hitsInformation?.hits) allResults = hitsInformation.hits;

  if (hitsInformation?.total?.value && hitsInformation.total.value <= store.pagination) {
    store.setHasMoreResults(false);
  }

  if (store.hasMoreResults && hitsInformation?.total?.value && allResults.length >= hitsInformation.total.value) {
    store.setIsEmailResponseForPagination(false);
  }

  if (!allResults || allResults?.length < 1) {
    store.setFetchedEmails([]);
    handleResetPagination();
    store.setIsEmailResponseForPagination(true);
    isLoading.value = false;
    notFound.value = true;
    return;
  }

  if (!isUserSearching && Array.isArray(store.fetchedEmails) && allResults.length > 0) {
    store.setFetchedEmails(allResults);
    if (hitsInformation?.total?.value)
      store.setIsEmailResponseForPagination(allResults.length < hitsInformation?.total?.value);
    isLoading.value = false;
    return;
  }

  if (Array.isArray(allResults) && allResults?.length > 0) {
    store.setFetchedEmails(allResults);
    if (hitsInformation?.total?.value)
      store.setIsEmailResponseForPagination(allResults.length < hitsInformation?.total?.value);
  }

  isLoading.value = false;
};

const handlePagination = () => {
  store.setPagination(PAGINATION_VALUE);
  handleSearch();
};

watch(searchValue, (element) => {
  if (!element) {
    store.setIsSearching(false);
    handleResetPagination();
    notFound.value = false;
  }
});

onMounted(() => {
  if (Array.isArray(store.fetchedEmails) && store.fetchedEmails.length < 1) handleSearch();
});

const goToEmail = () => {
  if (selectedEmail.value) {
    router.push({
      name: "email-content",
      params: { email_id: selectedEmail.value._source.message_id },
    });
  } else {
    alert("No hay email seleccionado");
  }
};
</script>

<template>
  <section class="p-6 max-w-4xl mx-auto">

    <div class="flex items-center space-x-3 mb-6">
      <img src="/Icons/MailIcon.png" alt="Logo" class="h-10 w-10">
      <h1 class="text-2xl font-bold">EmailSearch</h1>
    </div>

    <!-- Barra de búsqueda -->
    <SearchBar v-model:searchTerm="searchValue" v-model:SearchFunction="handleSearch" />
    <br>
    
    <div v-if="isLoading" class="flex justify-center py-6">
      <Loader />
    </div>

    <!-- Lista de correos -->
    <div v-if="hasResults" class="bg-white shadow-lg rounded-lg p-4">
      <ul class="divide-y divide-gray-200">
        <li 
          v-for="(email, index) in store.fetchedEmails" 
          :key="index" 
          @click="selectedEmail = email"
          class="p-4 cursor-pointer transition-all hover:bg-gray-50 rounded-lg"
        >
          <p class="text-sm text-gray-500">{{ email._source.date }}</p>
          <p class="text-lg font-semibold text-gray-800">{{ email._source.subject }}</p>
          <p class="text-sm text-gray-600">From: {{ email._source.from }}</p>
        </li>
      </ul>
    </div>

    <div v-else-if="notFound" class="flex justify-center py-6">
      <NotFound />
    </div>

    <!-- Botón de paginación -->
    <button 
      v-if="hasResults && store.hasMoreResults" 
      @click="handlePagination" 
      class="block mx-auto mt-6 px-5 py-2 bg-blue-600 text-white font-semibold rounded-lg transition hover:bg-blue-700"
    >
      Load more
    </button>

    <!-- Modal de detalles del email -->
    <div v-if="selectedEmail" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4" @click.self="selectedEmail = null">
      <div class="bg-white p-6 rounded-lg shadow-lg w-auto max-w-xl max-h-screen overflow-y-auto">
        <h2 class="text-lg font-bold mb-4">Email details</h2>
        <div class="grid grid-cols-1 gap-4">
          <div><strong>ID:</strong> {{ selectedEmail._id }}</div>
          <div><strong>From:</strong> {{ selectedEmail._source.from }}</div>
          <div><strong>To:</strong> {{ selectedEmail._source.to }}</div>
          <div><strong>Date:</strong> {{ selectedEmail._source.date }}</div>
          <div><strong>Subject:</strong> {{ selectedEmail._source.subject }}</div>
        </div>
        
        <div class="mt-4 flex justify-end gap-2">
          <button @click="selectedEmail = null" class="px-4 py-2 bg-red-600 text-white rounded">
            Close
          </button>
          <button @click="goToEmail" class="px-4 py-2 bg-blue-600 text-white rounded">
            See the full email
          </button>
        </div>
      </div>
    </div>
  </section>
</template>
