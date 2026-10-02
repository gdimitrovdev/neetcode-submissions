class Solution:
    def decodeString(self, s: str) -> str:
        ans = ""
        i = 0

        while i < len(s):
            num = ""
            
            while i < len(s) and s[i].isnumeric():
                num += s[i]
                i += 1

            if i == len(s): return ans

            if num != "" and s[i] == '[':
                open_br = 1
                i += 1
                num_int = int(num)
                substr = ""

                while i < len(s) and open_br != 0:
                    if s[i] == '[': open_br += 1
                    if s[i] == ']': open_br -= 1

                    if open_br > 0:
                        substr += s[i]

                    i += 1

                ans += self.decodeString(substr) * num_int

            else:
                ans += s[i]
                i += 1

        return ans
        