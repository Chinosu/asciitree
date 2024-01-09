// https://stackoverflow.com/questions/4965335/how-to-print-binary-tree-diagram-in-java


package alg;

import java.util.ArrayList;
import java.util.List;

class Main {
    public static void main(String[] args) {
        // Create nodes of the tree
        Node root = new Node(1);
        root.left = new Node(2);
        root.right = new Node(3);
        root.left.left = new Node(0);
        // root.left.right = new Node(5);
        // root.right.left = new Node(6);
        root.right.right = new Node(20);
        root.right.right.left = new Node(71);
        root.right.right.left.left = new Node(6);

        // Create an instance of TreeFormatter
        TreeFormatter formatter = new TreeFormatter();

        // Format and print the tree
        System.out.println(formatter.topDown(root));
    }
}

class Node {
    int data;
    Node left, right;

    Node(int data) {
        this.data = data;
        left = null;
        right = null;
    }
}


// ------------------------------------------------------------------------------------------------------------------------

class TreeFormatter {
    int padding = 2; // minimum number of horizontal spaces between two node data

    private int indent(List<String> lines, int margin) {
        // If negative, prefix all lines with spaces and return 0
        if (margin >= 0) return margin;
        String spaces = " ".repeat(-margin);
        int i = 0;
        for (var line : lines) {
            lines.set(i++, spaces + line);
        }
        return 0;
    }
    
    private List<String> merge(List<String> left, List<String> right) {
        // Merge two arrays, where the right strings are indented so there is no overlap
        int minSize = Math.min(left.size(), right.size());
        int offset = 0;
        for (int i = 0; i < minSize; i++) {
            offset = Math.max(offset, left.get(i).length() + padding - right.get(i).replaceAll("\\S.*", "").length());
        }
        indent(right, -indent(left, offset));
        for (int i = 0; i < minSize; i++) {
            left.set(i, left.get(i) + right.get(i).substring(left.get(i).length()));
        }
        if (right.size() > minSize) {
            left.addAll(right.subList(minSize, right.size()));
        }
        return left;
    }

    private List<String> buildLines(Node node) {
        if (node == null) return new ArrayList<>();
        List<String> lines = merge(buildLines(node.left), buildLines(node.right));
        int half = String.valueOf(node.data).length() / 2;
        int i = half;
        if (lines.size() > 0) {
            String line;
            i = lines.get(0).indexOf("*"); // Find index of first subtree
            if (node.right == null) {
                line = " ".repeat(i) + "┌─┘";
                i += 2;
            } else if (node.left == null) {
                line = " ".repeat(i = indent(lines, i - 2)) + "└─┐";
            } else {
                int dist = lines.get(0).length() - 1 - i; // Find distance between subtree roots
                line = String.format("%s┌%s┴%s┐", " ".repeat(i), "─".repeat(dist / 2 - 1), "─".repeat((dist - 1) / 2));
                i += dist / 2;
            }
            lines.set(0, line);
        }
        lines.add(0, " ".repeat(indent(lines, i - half)) + node.data);
        lines.add(0, " ".repeat(i + Math.max(0, half - i)) + "*"); // Add a marker for caller
        return lines;
    }
    
    public String topDown(Node root) {
        List<String> lines = buildLines(root);
        return String.join("\n", lines.subList(1, lines.size()));
    }
}