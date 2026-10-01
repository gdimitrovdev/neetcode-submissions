class Solution:
    def simplifyPath(self, path: str) -> str:
        curr = []

        for item in path.split('/'):
            if item == '' or item == '.':
                continue
            if item == '..':
                if len(curr) > 0:
                    curr.pop()
                continue

            curr.append(item)

        return '/' + '/'.join(curr)
        