function reverseWords(s: string): string {
    return s.split(" ").filter((item) => item !== "").reverse().join(" ")
};