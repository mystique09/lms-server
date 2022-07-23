<script>
  import Signup from "./lib/components/signup.svelte";
  import Dashboard from "./lib/components/dashboard/dashboard.svelte";
  import userStore from "./lib/stores/user";
</script>

<h1 class="font-bold text-lg">Hello, World!</h1>
{#if !$userStore}
  <Signup />
{:else}
  <Dashboard
    classrooms={[
      {
        description: "eksjej",
        name: "yo",
        subject: "yo",
        section: "kdkd",
        id: "jshshsh",
        room: "sjsj",
      },
    ]}
  />
{/if}
