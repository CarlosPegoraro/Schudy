<script>
    import axios from "axios";

    let username = "";
    let password = "";
    let showPassword = false;

    function handleSubmit() {
        console.log("Login attempt:", { username, password });
    }

    function togglePasswordVisibility() {
        showPassword = !showPassword;
    }
</script>

<div
    class="flex justify-center items-center min-h-screen p-4 bg-surface-50-900-token"
>
    <div class="card p-6 w-full max-w-md shadow-xl">
        <header class="card-header text-center pb-4">
            <h2 class="h2">Welcome Back</h2>
            <p class="text-surface-400-500-token">Sign in to your account</p>
        </header>

        <form on:submit|preventDefault={handleSubmit} class="space-y-4">
            <label class="label">
                <span class="label-text">Username</span>
                <input
                    bind:value={username}
                    type="text"
                    placeholder="Enter your email"
                    class="input input-bordered w-full"
                    required
                />
            </label>

            <label class="label">
                <span class="label-text">Password</span>
                <div
                    class="input-group input-group-divider grid-cols-[1fr_auto]"
                >
                    <input
                        bind:value={password}
                        type={showPassword ? "text" : "password"}
                        placeholder="Enter your password"
                        class="input input-bordered w-full"
                        required
                    />
                    <button
                        type="button"
                        class="variant-filled-surface"
                        on:click={togglePasswordVisibility}
                    >
                        {#if showPassword}
                            <span>Hide</span>
                        {:else}
                            <span>Show</span>
                        {/if}
                    </button>
                </div>
            </label>

            <div class="flex justify-between items-center">
                <label class="flex items-center space-x-2">
                    <input type="checkbox" class="checkbox" />
                    <span>Remember me</span>
                </label>
                <!-- svelte-ignore a11y_invalid_attribute -->
                <a href="#" class="anchor text-sm">Forgot password?</a>
            </div>

            <button
                type="submit"
                class="btn btn-variant-filled bg-secondary-500 w-full"
            >
                <span>Login</span>
            </button>
        </form>

        <div class="text-center pt-4">
            <p>
                Don't have an account? <a href="/register" class="anchor"
                    >Sign up</a
                >
            </p>
        </div>
    </div>
</div>
