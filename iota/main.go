package main

import (
	"fmt"
	"os"
)

const (
	Read    = 1 << iota // 1 << 0 = 1
	Write               // 1 << 1 = 2
	Execute             // 1 << 2 = 4
)

type Role struct {
	Visitors, Readers, Writers, Admins uint8
}

type File struct {
	Permissions uint8
}

func initializeRoles() Role {
	return Role{
		Visitors: 0,
		Readers:  Read,
		Writers:  Read | Write,
		Admins:   Read | Write | Execute,
	}
}

func hasPermission(permission, permToCheck uint8) bool {
	return permission&permToCheck == permToCheck
}

func createReadOnlyFile() File {
	return File{
		Permissions: Read,
	}
}

func main() {
	roles := initializeRoles()

	fmt.Println("Roles and their Permissions:")
	fmt.Println("Visitors:", roles.Visitors)
	fmt.Println("Readers:", roles.Readers)
	fmt.Println("Writers:", roles.Writers)
	fmt.Println("Admins:", roles.Admins)

	fmt.Println("\nVisitor permissions:")
	fmt.Println("Has Read Access?", hasPermission(roles.Visitors, Read))
	fmt.Println("Has Write Access?", hasPermission(roles.Visitors, Write))
	fmt.Println("Has Execute Access?", hasPermission(roles.Visitors, Execute))

	fmt.Println("\nReader permissions:")
	fmt.Println("Has Read Access?", hasPermission(roles.Readers, Read))
	fmt.Println("Has Write Access?", hasPermission(roles.Readers, Write))
	fmt.Println("Has Execute Access?", hasPermission(roles.Readers, Execute))

	fmt.Println("\nWriter permissions:")
	fmt.Println("Has Read Access?", hasPermission(roles.Writers, Read))
	fmt.Println("Has Write Access?", hasPermission(roles.Writers, Write))
	fmt.Println("Has Execute Access?", hasPermission(roles.Writers, Execute))

	fmt.Println("\nAdmin permissions:")
	fmt.Println("Has Read Access?", hasPermission(roles.Admins, Read))
	fmt.Println("Has Write Access?", hasPermission(roles.Admins, Write))
	fmt.Println("Has Execute Access?", hasPermission(roles.Admins, Execute))

	// Crear un archivo con permiso de solo lectura
	readOnlyFile := createReadOnlyFile()
	fmt.Println("\nFile Permissions:", readOnlyFile.Permissions)

	// Crear un archivo con permisos de solo lectura
	fileName := "example.txt"
	err := os.WriteFile("./iota/"+fileName, []byte("Este es un archivo de solo lectura"), 0400)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}

	// Verificar los permisos del archivo recién creado
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("Error al obtener información del archivo:", err)
		return
	}

	// Imprimir los permisos del archivo
	fmt.Printf("Permisos del archivo %s: %o\n", fileName, fileInfo.Mode().Perm())
}
