<script lang="ts">
    // Define types for form data
    interface FormData {
        username: string;
        email: string;
        confirmEmail: string;
        password: string;
        confirmPassword: string;
        dateOfBirth: string;
    }

    // Define types for form errors
    interface FormErrors {
        username?: string;
        email?: string;
        confirmEmail?: string;
        password?: string;
        confirmPassword?: string;
        dateOfBirth?: string;
        [key: string]: string | undefined;
    }

    // Form data with initial values
    let formData: FormData = {
        username: "",
        email: "",
        confirmEmail: "",
        password: "",
        confirmPassword: "",
        dateOfBirth: "",
    };

    // Form validation state
    let errors: FormErrors = {};
    let showPassword: boolean = false;
    let showConfirmPassword: boolean = false;

    // Toggle password visibility
    function togglePassword(): void {
        showPassword = !showPassword;
    }

    function toggleConfirmPassword(): void {
        showConfirmPassword = !showConfirmPassword;
    }

    // Form submission
    function handleSubmit(event: SubmitEvent): void {
        // Prevent default form submission
        event.preventDefault();

        // Reset errors
        errors = {};

        // Basic validation
        if (!formData.username) errors.username = "Username is required";
        if (!formData.email) errors.email = "Email is required";
        if (formData.email && !isValidEmail(formData.email))
            errors.email = "Please enter a valid email";
        if (formData.email !== formData.confirmEmail)
            errors.confirmEmail = "Emails do not match";
        if (!formData.password) errors.password = "Password is required";
        if (formData.password && formData.password.length < 8)
            errors.password = "Password must be at least 8 characters";
        if (formData.password !== formData.confirmPassword)
            errors.confirmPassword = "Passwords do not match";
        if (!formData.dateOfBirth)
            errors.dateOfBirth = "Date of birth is required";

        // If no errors, submit the form
        if (Object.keys(errors).length === 0) {
            console.log("Form submitted:", formData);
            // Here you would typically submit to your backend
        }
    }

    // Email validation helper
    function isValidEmail(email: string): boolean {
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailRegex.test(email);
    }

    // Password strength check
    function getPasswordStrength(password: string): {
        strength: number;
        message: string;
    } {
        if (!password) return { strength: 0, message: "No password" };

        let strength = 0;
        if (password.length >= 8) strength += 1;
        if (/[A-Z]/.test(password)) strength += 1;
        if (/[a-z]/.test(password)) strength += 1;
        if (/[0-9]/.test(password)) strength += 1;
        if (/[^A-Za-z0-9]/.test(password)) strength += 1;

        let message = "";
        switch (strength) {
            case 0:
            case 1:
                message = "Weak";
                break;
            case 2:
            case 3:
                message = "Medium";
                break;
            case 4:
            case 5:
                message = "Strong";
                break;
            default:
                message = "";
        }

        return { strength, message };
    }

    // Reactive password strength
    $: passwordStrength = getPasswordStrength(formData.password);
</script>

<div
    class="flex justify-center items-center min-h-screen p-4 bg-surface-50-900-token"
>
    <div class="card py-6 w-full max-w-md shadow-xl">
        <header class="card-header text-center pb-4">
            <h2 class="h2">Create Account</h2>
            <p class="text-surface-400-500-token">Join our community today</p>
        </header>

        <form on:submit={handleSubmit} class="space-y-4">
            <!-- Username -->
            <label class="label">
                <span class="label-text">Username</span>
                <input
                    bind:value={formData.username}
                    type="text"
                    placeholder="Choose a username"
                    class="input input-bordered w-full {errors.username
                        ? 'input-error'
                        : ''}"
                />
                {#if errors.username}
                    <span class="text-error-500">{errors.username}</span>
                {/if}
            </label>

            <!-- Email -->
            <label class="label">
                <span class="label-text">Email</span>
                <input
                    bind:value={formData.email}
                    type="email"
                    placeholder="Enter your email"
                    class="input input-bordered w-full {errors.email
                        ? 'input-error'
                        : ''}"
                />
                {#if errors.email}
                    <span class="text-error-500">{errors.email}</span>
                {/if}
            </label>

            <!-- Confirm Email -->
            <label class="label">
                <span class="label-text">Confirm Email</span>
                <input
                    bind:value={formData.confirmEmail}
                    type="email"
                    placeholder="Confirm your email"
                    class="input input-bordered w-full {errors.confirmEmail
                        ? 'input-error'
                        : ''}"
                />
                {#if errors.confirmEmail}
                    <span class="text-error-500">{errors.confirmEmail}</span>
                {/if}
            </label>

            <!-- Password -->
            <label class="label">
                <span class="label-text">Password</span>
                <div
                    class="input-group input-group-divider grid-cols-[1fr_auto]"
                >
                    <input
                        bind:value={formData.password}
                        type={showPassword ? "text" : "password"}
                        placeholder="Create a password"
                        class="input input-bordered w-full {errors.password
                            ? 'input-error'
                            : ''}"
                    />
                    <button
                        type="button"
                        class="variant-filled-surface"
                        on:click={togglePassword}
                    >
                        {#if showPassword}
                            <span>Hide</span>
                        {:else}
                            <span>Show</span>
                        {/if}
                    </button>
                </div>
                {#if formData.password}
                    <div class="mt-1">
                        <div
                            class="w-full bg-surface-200-700-token rounded-full h-1.5"
                        >
                            <div
                                class="h-1.5 rounded-full {passwordStrength.strength <=
                                2
                                    ? 'bg-error-500'
                                    : passwordStrength.strength <= 3
                                      ? 'bg-warning-500'
                                      : 'bg-success-500'}"
                                style="width: {passwordStrength.strength * 20}%"
                            ></div>
                        </div>
                        <span class="text-xs">{passwordStrength.message}</span>
                    </div>
                {/if}
                {#if errors.password}
                    <span class="text-error-500">{errors.password}</span>
                {/if}
            </label>

            <!-- Confirm Password -->
            <label class="label">
                <span class="label-text">Confirm Password</span>
                <div
                    class="input-group input-group-divider grid-cols-[1fr_auto]"
                >
                    <input
                        bind:value={formData.confirmPassword}
                        type={showConfirmPassword ? "text" : "password"}
                        placeholder="Confirm your password"
                        class="input input-bordered w-full {errors.confirmPassword
                            ? 'input-error'
                            : ''}"
                    />
                    <button
                        type="button"
                        class="variant-filled-surface"
                        on:click={toggleConfirmPassword}
                    >
                        {#if showConfirmPassword}
                            <span>Hide</span>
                        {:else}
                            <span>Show</span>
                        {/if}
                    </button>
                </div>
                {#if errors.confirmPassword}
                    <span class="text-error-500">{errors.confirmPassword}</span>
                {/if}
            </label>

            <!-- Date of Birth - Using standard HTML date input -->
            <label class="label">
                <span class="label-text">Date of Birth</span>
                <input
                    bind:value={formData.dateOfBirth}
                    type="date"
                    class="input input-bordered w-full {errors.dateOfBirth
                        ? 'input-error'
                        : ''}"
                />
                {#if errors.dateOfBirth}
                    <span class="text-error-500">{errors.dateOfBirth}</span>
                {/if}
            </label>

            <!-- Terms and Conditions -->
            <label class="flex items-center space-x-2">
                <input type="checkbox" class="checkbox" />
                <!-- svelte-ignore a11y_invalid_attribute -->
                <span
                    >I agree to the <a href="#" class="anchor"
                        >Terms and Conditions</a
                    ></span
                >
            </label>

            <!-- Submit Button -->
            <button
                type="submit"
                class="btn variant-filled-primary w-full bg-secondary-500"
            >
                <span>Create Account</span>
            </button>
        </form>

        <div class="text-center pt-4">
            <p>
                Already have an account? <a
                    href="/login"
                    class="anchor text-primary-500">Sign in</a
                >
            </p>
        </div>
    </div>
</div>
