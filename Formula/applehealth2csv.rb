# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Applehealth2csv < Formula
  desc "Command line tool to convert Apple Watch Health Data to CSV or JSON files"
  homepage "https://github.com/muquit/applehealth2csv"
  version "1.0.1"
  license "MIT"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/muquit/applehealth2csv/releases/download/v1.0.1/applehealth2csv_1.0.1_mac-64bit.tar.gz"
    sha256 "7771463fcc5f90293b02cf7983153f9d8eb50e603a3079fe2ba614361147f5cb"
  end
  if OS.linux? && Hardware::CPU.intel?
    url "https://github.com/muquit/applehealth2csv/releases/download/v1.0.1/applehealth2csv_1.0.1_linux-64bit.tar.gz"
    sha256 "3a0c8e91e205540dd128e6502988eeb361bdc0b033fa30242188700c640c2c10"
  end

  def install
    bin.install "applehealth2csv"
    man1.install "docs/applehealth2csv.1"
  end
end