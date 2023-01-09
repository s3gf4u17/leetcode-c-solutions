/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func preorderTraversal(root *TreeNode) []int {
    response:=[]int{}
    if root!=nil{
        response=append(response,root.Val)
        if root.Left!=nil {
            response=append(response,preorderTraversal(root.Left)...)
        }
        response=append(response,preorderTraversal(root.Right)...)
    }
    return response
}