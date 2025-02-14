package main

import (
	"fmt"
	"log"

	"github.com/orbit-center/sdk/client"
)

func main() {
	// 创建客户端
	c := client.NewClient("http://localhost:8081/api/v1", "your-token")

	// 测试用户服务
	testUserService(c)

	// 测试角色服务
	testRoleService(c)

	// 测试菜单服务
	testMenuService(c)

	// 测试字典服务
	testDictService(c)
}

func testUserService(c *client.Client) {
	user, err := c.User().GetUserInfo()
	if err != nil {
		log.Printf("获取用户信息失败: %v", err)
		return
	}
	fmt.Printf("当前用户: %+v\n", user)

	// 测试权限检查
	hasPermission := c.User().CheckPermission("system:user:view")
	fmt.Printf("是否有用户查看权限: %v\n", hasPermission)
}

func testRoleService(c *client.Client) {
	roles, err := c.Role().GetRoleList()
	if err != nil {
		log.Printf("获取角色列表失败: %v", err)
		return
	}
	fmt.Printf("角色列表: %+v\n", roles)

	if len(roles) > 0 {
		permissions, err := c.Role().GetRolePermissions(roles[0].ID)
		if err != nil {
			log.Printf("获取角色权限失败: %v", err)
			return
		}
		fmt.Printf("角色权限: %+v\n", permissions)
	}
}

func testMenuService(c *client.Client) {
	menus, err := c.Menu().GetMenuList()
	if err != nil {
		log.Printf("获取菜单列表失败: %v", err)
		return
	}
	fmt.Printf("菜单列表: %+v\n", menus)

	if len(menus) > 0 {
		hasPermission := c.Menu().CheckMenuPermission(menus[0].ID)
		fmt.Printf("是否有菜单访问权限: %v\n", hasPermission)
	}
}

func testDictService(c *client.Client) {
	dicts, err := c.Dict().GetDictList()
	if err != nil {
		log.Printf("获取字典列表失败: %v", err)
		return
	}
	fmt.Printf("字典列表: %+v\n", dicts)

	if len(dicts) > 0 {
		types, err := c.Dict().GetDictTypes(dicts[0].ID)
		if err != nil {
			log.Printf("获取字典类型失败: %v", err)
			return
		}
		fmt.Printf("字典类型: %+v\n", types)

		if len(types) > 0 {
			contents, err := c.Dict().GetDictContents(types[0].ID)
			if err != nil {
				log.Printf("获取字典内容失败: %v", err)
				return
			}
			fmt.Printf("字典内容: %+v\n", contents)
		}
	}
}
