package com.feyfree.leetcode.p0130;

import java.util.HashMap;
import java.util.Map;

/**
 * 130. 被围绕的区域
 * https://leetcode-cn.com/problems/surrounded-regions/
 *
 * @author leilei
 */
class Solution {

    private final int[] dx = {-1, 1, 0, 0};
    private final int[] dy = {0, 0, -1, 1};

    public void solve(char[][] board) {
        int m = board.length;
        if (m == 0) {
            return;
        }
        int n = board[0].length;
        for (int y = 0; y < m; ++y) {
            dfs(0, y, board, m, n);
            dfs(n - 1, y, board, m, n);
        }
        for (int x = 0; x < n; ++x) {
            dfs(x, 0, board, m, n);
            dfs(x, m - 1, board, m, n);
        }
        Map<Character, Character> map = new HashMap<>();
        map.put('G', 'O');
        map.put('O', 'X');
        map.put('X', 'X');
        for (int y = 0; y < m; ++y) {
            for (int x = 0; x < n; ++x) {
                board[y][x] = map.get(board[y][x]);
            }
        }

    }

    private void dfs(int x, int y, char[][] board, int m, int n) {
        if (x < 0 || x >= n || y < 0 || y >= m || board[y][x] != 'O') {
            return;
        }
        // mark it as good
        board[y][x] = 'G';
        for (int i = 0; i < dx.length; i++) {
            dfs(x + dx[i], y + dy[i], board, m, n);
        }
    }
}