package cmd

import (
	"testing"
)

func TestBuildMkdocUrl_ReturnsCorrectUrlForSimplePath(t *testing.T) {
	baseFolder := "/home/user/docs"
	currentFolder := "/home/user/docs/section/page.md"
	baseUrl := "https://example.com/"
	expected := "https://example.com/section/page/"

	result := buildMkdocUrl(baseFolder, currentFolder, baseUrl)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestBuildMkdocUrl_BaseUrlWithTrailingSlash(t *testing.T) {
	baseFolder := "/docs"
	currentFolder := "/docs/guide/intro.md"
	baseUrl := "https://site.org/"
	expected := "https://site.org/guide/intro/"

	result := buildMkdocUrl(baseFolder, currentFolder, baseUrl)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestBuildMkdocUrl_CurrentFolderIsBaseFolder(t *testing.T) {
	baseFolder := "/docs"
	currentFolder := "/docs/index.md"
	baseUrl := "https://site.org/docs"
	expected := "https://site.org/docs/index/"

	result := buildMkdocUrl(baseFolder, currentFolder, baseUrl)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestBuildMkdocUrl_BaseUrlWithoutTrailingSlash(t *testing.T) {
	baseFolder := "/docs"
	currentFolder := "/docs/guide/intro.md"
	baseUrl := "https://site.org/docs"
	expected := "https://site.org/docs/guide/intro/"

	result := buildMkdocUrl(baseFolder, currentFolder, baseUrl)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestBuildMkdocUrl_ReturnsBaseUrlOnError(t *testing.T) {
	baseFolder := "/docs"
	currentFolder := "/other/guide/intro.md"
	baseUrl := "https://site.org/docs"
	expected := baseUrl

	result := buildMkdocUrl(baseFolder, currentFolder, baseUrl)
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}
