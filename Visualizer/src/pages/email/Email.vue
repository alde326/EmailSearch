<script setup lang="ts">
import { ref, onMounted, watch } from "vue";
import { useRoute } from "vue-router";

import Loader from "../../components/Loader.vue";
import NotFound from "../../components/NotFound.vue";

import {
  getEmailFooterDetails,
  getParsedEmailId,
  getEmailLastHeaderDetails,
} from "../../helpers/emailDetails";

import getSingleEmail from "../../helpers/fetchSingleEmail";

import { Source } from "../../types/emailTypes";

import store from "../../Store";

window.scrollTo(0, 0);

const route = useRoute();
const parsedEmailId: string | null | undefined = getParsedEmailId(route);

const emailSheetRef = ref<HTMLElement | null>(null);
const isLoadingRef = ref<boolean>(true);
const isErrorRef = ref<boolean>(false);

const allEmailContent = ref<Source | null | undefined>();
const timestamp = ref<string | null | undefined>();
const date = ref<string | null | undefined>();
const sent = ref<string | null | undefined>();
const dateSub = ref<string | null | undefined>();
const subject = ref<string | null | undefined>();
const body = ref<string | null | undefined>();
const cc = ref<string | null | undefined>();
const xCc = ref<string | null | undefined>();

watch([allEmailContent, emailSheetRef], () => {
  const handleHighLightText = () => {
    if (!emailSheetRef.value || !store?.searchedValue || !allEmailContent.value)
      return;
    const allFields = emailSheetRef.value.querySelectorAll("p");
    if (!allFields) return;
    for (const field of allFields) {
      let emailContent = field.innerHTML;
      const regex = new RegExp(`(${store.searchedValue})`, "gi");
      if (regex.test(emailContent)) {
        emailContent = emailContent.replace(
          regex,
          `<mark class="highlight">${store.searchedValue}</mark>`
        );
      }
      field.innerHTML = emailContent;
    }
  };
  handleHighLightText();
});

onMounted(() => {
  const loadEmailProps = () => {
    if (!allEmailContent.value) {
      isErrorRef.value = true;
      return;
    }
    timestamp.value = allEmailContent.value?.["@timestamp"];
    date.value = allEmailContent.value?.date;
    sent.value = allEmailContent.value?.sent;
    dateSub.value = allEmailContent.value?.date_subemail;
    subject.value = allEmailContent.value?.subject;
    body.value = allEmailContent.value?.body;
    cc.value = allEmailContent.value?.cc;
    xCc.value = allEmailContent.value?.x_cc;
  };

  const getEmail = async () => {
    if (!parsedEmailId) return;
    isLoadingRef.value = true;
    allEmailContent.value = await getSingleEmail(parsedEmailId);
    loadEmailProps();
    isLoadingRef.value = false;
  };
  getEmail();
});
</script>
<template>
  <section
    v-if="!isLoadingRef && !isErrorRef"
    ref="emailSheetRef"
    class="email-sheet max-w-[800px] mx-auto mt-16 p-6 bg-white shadow-md rounded-lg grid grid-cols-1 gap-6"
  >

    <div class="flex items-center space-x-3 mb-6">
      <img src="/Icons/MailIcon.png" alt="Logo" class="h-10 w-10">
      <h1 class="text-2xl font-bold">EmailSearch</h1>
    </div>

    <header class="col-span-full mb-6">
      <h2 class="text-xl font-semibold text-gray-800">Email Details</h2>
    </header>

    <!-- ID del Correo -->
    <section class="grid grid-cols-2 gap-4 mb-6">
      <h2 class="text-lg font-medium text-gray-700 col-span-2">Email Date</h2>
      <p class="text-sm text-gray-600 col-span-2">
        {{ timestamp ? new Date(timestamp).toLocaleString() : "No disponible" }}
      </p>
    </section>

    <!-- Encabezados: From, To, etc. -->
    <section class="grid grid-cols-1 gap-4 mb-6">
      <h2 class="text-lg font-medium text-gray-700">Headers</h2>
      <div v-for="(header, index) in getEmailLastHeaderDetails(allEmailContent)" :key="index" class="grid grid-cols-2 gap-4">
        <p class="font-semibold text-gray-700">{{ header.detailName }}:</p>
        <p class="text-sm text-gray-600">{{ header.detailFirstValue }} <span v-if="header.detailSecondValue"> / {{ header.detailSecondValue }}</span></p>
      </div>
    </section>

    <!-- Asunto -->
    <section class="grid grid-cols-2 gap-4 mb-6">
      <h2 class="text-lg font-medium text-gray-700 col-span-2">Subject</h2>
      <p class="text-sm text-gray-600 col-span-2">{{ subject }}</p>
    </section>

    <!-- Cuerpo del Correo -->
    <section class="grid grid-cols-1 gap-4 mb-6">
      <h2 class="text-lg font-medium text-gray-700">Email body</h2>
      <div v-html="body" class="text-gray-700 text-base"></div>
    </section>

    <!-- Detalles adicionales -->
    <footer class="mt-6 text-sm text-gray-500 grid grid-cols-1 gap-4">
      <div v-if="cc || xCc" class="mb-2">
        <p class="font-medium">CC:</p>
        <p>{{ cc || xCc }}</p>
      </div>

      <div v-for="(details, _) in getEmailFooterDetails(allEmailContent)" :key="details.detailName">
        <div v-if="details?.detailValue">
          <p class="font-medium">{{ details.detailName }}:</p>
          <p>{{ details.detailValue }}</p>
        </div>
      </div>
    </footer>
  </section>

  <div
    v-else-if="isLoadingRef && !isErrorRef"
    class="flex items-center justify-center h-screen w-full bg-gray-100"
  >
    <Loader />
  </div>

  <div
    v-else-if="isErrorRef"
    class="flex items-center justify-center h-screen w-full bg-gray-100"
  >
    <NotFound :isEmailError="isErrorRef" />
  </div>
</template>

<style scoped>
.email-sheet {
  box-shadow: 0px 2px 6px rgba(0, 0, 0, 0.1);
}

.email-sheet p {
  overflow-wrap: break-word; /* Asegura que el texto largo se ajuste dentro del contenedor */
}

.email-sheet section h2 {
  margin-bottom: 4px;
  color: #4b5563;
}

.email-body p {
  line-height: 1.6;
}

.email-footer p {
  margin-bottom: 4px;
}

.highlight {
  background-color: #ffeb3b;
  color: #333;
}
</style>