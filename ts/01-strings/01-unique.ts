// IsUnique results if a string has all unique chars
// This solution is using HashMap data struct
// Big O (N)
export function isUnique(input: string) {
    const hashMap = new Map<string, boolean>();

    for (let i = 0; i < input.length; i++) {
        if (hashMap.has(input[i])) {
            return false;
        }

        hashMap.set(input[i], true);
    }

    return hashMap.size === input.length;
}

// IsUniqueSimple results if a string has all unique chars
// This solution is using none data struct, just two for loops
// Big O (N^2)
export function isUniqueSimple(input: string) {
    for (let i = 0; i < input.length; i++) {
        for (let j = i + 1; j < input.length; j++) {
            if (input[i] === input[j]) {
                return false;
            }
        }
    }

    return true;
}

// IsUniqueSort results if a string has all unique chars
// This solution is using none data struct, but sorting the string
// Big O (n log(n))
export function isUniqueSort(input: string) {
    const sorted = input.split("").sort().join("");

    for (let i = 0; i < sorted.length; i++) {
        if (sorted[i] === sorted[i + 1]) {
            return false;
        }
    }

    return true;
}
